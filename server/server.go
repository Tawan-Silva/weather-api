package server

import (
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
	httpSwagger "github.com/swaggo/http-swagger"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	_ "zip_temperature/docs" // Importa os documentos gerados pelo swag
	"zip_temperature/handlers"
)

func Start() {

	router := mux.NewRouter()
	router.Use(CorsMiddleware)
	apiRouter := router.PathPrefix("/api-weather").Subrouter()
	apiRouter.HandleFunc("/health", healthCheckHandler)
	apiRouter.HandleFunc("/get", handlers.ZipCodeHandler).Methods(http.MethodGet)
	apiRouter.PathPrefix("/swagger-ui/").Handler(httpSwagger.WrapHandler)
	apiRouter.HandleFunc("/swagger.json", GetSwaggerFile)

	log.Println("Starting server on :8080")
	go browser.OpenURL("api-weather/swagger-ui/index.html")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func GetSwaggerFile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./docs/swagger.json")
	if err != nil {
		log.Printf("Failed to read Swagger JSON file: %v", err)
		http.Error(w, "Failed to read Swagger JSON file", http.StatusInternalServerError)
		return
	}

	host := os.Getenv("HOST")

	modifiedData := strings.ReplaceAll(string(data), "localhost:8080", host)

	w.Write([]byte(modifiedData))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API OK"))
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
