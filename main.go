package main

import (
	"log"
	"zip_temperature/server"
)

// @title Weather API
// @version 1.0
// @description This is a simple weather API
// @termsOfService http://swagger.io/terms/
//
// @contact.name Tawan Silva
// @contact.url https://www.linkedin.com/in/tawan-silva-684b581b7/
// @contact.email tawan.tls43@gmail.com
//
// @license.name Weather License
// @license.url http://www.weather.com.br
//
// @host localhost:8080
// @BasePath /api-weather
func main() {
	server.Start()
	log.Println("Server started successfully")
}
