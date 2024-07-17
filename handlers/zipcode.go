package handlers

import (
	"encoding/json"
	"net/http"
	"zip_temperature/services"
)

// ErrorResponse godoc
type ErrorResponse struct {
	Message string `json:"message"`
}

// ZipCodeHandler godoc
// @Summary Get weather from zipcode
// @Description Get weather from zipcode
// @Tags zipcode
// @Accept json
// @Produce json
// @Param zipcode query string true "Zipcode"
// @Success 200 {object} models.WeatherResponse
// @Failure 404 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Router /get [get]
func ZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	zipcode := r.URL.Query().Get("zipcode")
	if len(zipcode) != 8 {
		responseError := ErrorResponse{Message: "invalid zipcode"}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(responseError)
		return
	}

	lat, lon, err := services.FetchLatLonFromZipCode(zipcode)
	if err != nil {
		responseError := ErrorResponse{Message: "can not find zipcode"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(responseError)
		return
	}

	weather, err := services.FetchWeather(lat, lon)
	if err != nil {
		responseError := ErrorResponse{Message: "can not find weather"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(responseError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}
