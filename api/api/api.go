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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	town := r.URL.Query().Get("town")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	offsetInt, offsetErr := strconv.Atoi(offset)
	limitInt, limitErr := strconv.Atoi(limit)

	dateRangeStart, startErr := time.Parse("2006-01-02T15:04:05.999Z", from)
	dateRangeEnd, endErr := time.Parse("2006-01-02T15:04:05.999Z", to)

	filtered := Filter(WeatherData, func(w Weather) bool {
		filter := true
		if town != "" {
			filter = strings.ToLower(town) == strings.ToLower(w.Town)
		}
		if filter == false {
			return filter
		}
		if startErr == nil || endErr == nil {
			start, sErr := time.Parse(DateFormat, w.Date)
			end, eErr := time.Parse(DateFormat, w.Date)

			if startErr == nil && sErr == nil {
				filter = start.After(dateRangeStart) || start.Equal(dateRangeStart)
				if filter == false {
					return filter
				}
			}
			if endErr == nil && eErr == nil {
				filter = end.Before(dateRangeEnd) || end.Equal(dateRangeEnd)
			}
		}

		return filter
	})

	length := len(filtered)

	if offsetErr == nil && offsetInt >= 0 {
		if offsetInt > len(filtered) {
			filtered = []Weather{}
		} else {
			filtered = filtered[offsetInt:len(filtered)]
		}
	}

	if limitErr == nil && limitInt >= 0 {
		if limitInt < len(filtered) {
			filtered = filtered[0:limitInt]
		}
	}

	response := WeatherResponse{
		Data:   filtered,
		Limit:  limitInt,
		Offset: offsetInt,
		Count:  length,
	}

	json.NewEncoder(w).Encode(response)
}

func Filter(it []Weather, f func(Weather) bool) []Weather {
	filtered := make([]Weather, 0)
	for _, v := range it {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func RequestHandler() {
	http.HandleFunc("/api/weather", fetchWeather)
	http.HandleFunc("/api/healthz", healthz)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
