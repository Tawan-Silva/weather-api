package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zip_temperature/configs"
	"zip_temperature/models"
)

type GeoapifyResponse struct {
	Features []struct {
		Properties struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"properties"`
	} `json:"features"`
}

func FetchLatLonFromZipCode(zipcode string) (string, string, error) {
	localtion, err := FetchCityFromZipCode(zipcode)
	if err != nil {
		return "", "", err
	}
	key := configs.GetGeoApiKey()
	resp, err := http.Get(fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", localtion, key))
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var data GeoapifyResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", "", err
	}

	if len(data.Features) == 0 {
		return "", "", fmt.Errorf("zipcode not found")
	}

	lat := data.Features[0].Properties.Lat
	lon := data.Features[0].Properties.Lon

	return fmt.Sprintf("%.6f", lat), fmt.Sprintf("%.6f", lon), nil
}

func FetchCityFromZipCode(zipcode string) (string, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + zipcode + "/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data models.ZipCodeResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", fmt.Errorf("zipcode not found")
	}

	return fmt.Sprintf("%s, %s, %s, %s", data.Logradouro, data.Bairro, data.Localidade, data.UF), nil
}
