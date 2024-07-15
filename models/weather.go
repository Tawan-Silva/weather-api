package models

type WeatherResponse struct {
	TemperatureC string `json:"temp_C"`
	TemperatureF string `json:"temp_F"`
	TemperatureK string `json:"temp_K"`
}
