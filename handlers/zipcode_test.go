package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zip_temperature/models"
)

func TestZipCodeHandler(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ZipCodeHandler(w, r)
	})

	t.Run("returns weather for valid zipcode", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/get?zipcode=35790006", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var weather models.WeatherResponse
		err := json.NewDecoder(rr.Body).Decode(&weather)
		if err != nil {
			t.Errorf("handler returned unexpected error: %v", err)
		}
	})

	t.Run("returns error for invalid zipcode", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/get?zipcode=123", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
		}
	})

	t.Run("returns error when zipcode not found", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/get?zipcode=00000000", nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}
