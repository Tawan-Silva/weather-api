package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zip_temperature/models"
)

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

	return data.Localidade, nil
}
