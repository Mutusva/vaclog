package main

import (
	appConfig "Codenotary/config"
	"Codenotary/internal/api"
	"Codenotary/internal/immuDB"
	"Codenotary/internal/router"
	"net/http"
)

// @title Vaccination log for animals
// @version 2.0
// @description This API provides functionality to record vaccination history for farm animal
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	config := appConfig.ReadConfig()
	client := &http.Client{}
	immuDBClient := immuDB.New(config, client)

	controller := api.NewVacController(config, immuDBClient)
	router.Init(controller)
}
