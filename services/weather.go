package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zip_temperature/configs"
	"zip_temperature/models"
	"zip_temperature/utils"
)

type WeatherAPIResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func FetchWeather(lat string, lon string) (models.WeatherResponse, error) {
	key := configs.GetWeatherApiKey()
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s,%s", key, lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return models.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var data WeatherAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	tempC := data.Current.TempC
	tempF := utils.CelsiusToFahrenheit(tempC)
	tempK := utils.CelsiusToKelvin(tempC)

	return models.WeatherResponse{
		TemperatureC: fmt.Sprintf("%.2f", tempC),
		TemperatureF: fmt.Sprintf("%.2f", tempF),
		TemperatureK: fmt.Sprintf("%.2f", tempK),
	}, nil
}
