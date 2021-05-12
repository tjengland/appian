package api

import (
	"log"
	"net/http"
)

type WeatherResponse struct {
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Data   []Weather `json:"data"`
	Count  int       `json:"count"`
}

func fetchWeather(w http.ResponseWriter, r *http.Request) {
	town := r.URL.Query().Get("town")
	dateRange := r.URL.Query().Get("date")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	filtered := Filter(WeatherData, func(w Weather) bool {
		filter := true
		if location != nil {
			filter = location == w.Location
		}
		if filer == false {
			return filter
		}
		if dateRange != nil && (dateRange[0] && dateRange[1]) {

		}

		return filter
	})

	response := WeatherResponse{
		Data:   filtered,
		Limit:  limit,
		Offset: Offset,
		Count:  len(filtered),
	}
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

func RequestHandler() {
	http.HandleFunc("/api/weather", fetchWeather)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
