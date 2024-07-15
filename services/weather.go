package services

import (
	"encoding/json"
	"net/http"
	"zip_temperature/configs"
	"zip_temperature/models"
	"zip_temperature/utils"
)

func FetchWeather(city string) (models.WeatherResponse, error) {
	key := configs.GetWeatherApiKey()
	resp, err := http.Get("https://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + city)
	if err != nil {
		return models.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	tempC := data.Current.TempC
	tempF := utils.CelsiusToFahrenheit(tempC)
	tempK := utils.CelsiusToKelvin(tempC)

	return models.WeatherResponse{
		TemperatureC: tempC,
		TemperatureF: tempF,
		TemperatureK: tempK,
	}, nil
}
