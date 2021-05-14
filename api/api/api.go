package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type WeatherResponse struct {
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Data   []Weather `json:"data"`
	Count  int       `json:"count"`
}

func fetchWeather(w http.ResponseWriter, r *http.Request) {
	town := r.URL.Query().Get("town")
	offset, offsetErr := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, limitErr := strconv.Atoi(r.URL.Query().Get("limit"))
	from, fromErr := time.Parse("2006-01-02T15:04:05.999Z", r.URL.Query().Get("from"))
	to, toErr := time.Parse("2006-01-02T15:04:05.999Z", r.URL.Query().Get("to"))

	filtered := Filter(WeatherData, func(w Weather) bool {
		filter := true
		if town != "" {
			filter = strings.ToLower(town) == strings.ToLower(w.Town)
		}
		if filter == false {
			return filter
		}
		if fromErr == nil || toErr == nil {
			start, sErr := time.Parse(DateFormat, w.Date)
			end, eErr := time.Parse(DateFormat, w.Date)

			if fromErr == nil && sErr == nil {
				filter = start.After(from) || start.Equal(from)
				if filter == false {
					return filter
				}
			}
			if toErr == nil && eErr == nil {
				filter = end.Before(to) || end.Equal(to)
			}
		}

		return filter
	})

	length := len(filtered)

	if offsetErr == nil && offset >= 0 {
		if offset > len(filtered) {
			filtered = []Weather{}
		} else {
			filtered = filtered[offset:len(filtered)]
		}
	}

	if limitErr == nil && limit >= 0 {
		if limit < len(filtered) {
			filtered = filtered[0:limit]
		}
	}

	response := WeatherResponse{
		Data:   filtered,
		Limit:  limit,
		Offset: offset,
		Count:  length,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func RequestHandler() {
	http.HandleFunc("/api/weather", fetchWeather)
	http.HandleFunc("/api/healthz", healthz)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
