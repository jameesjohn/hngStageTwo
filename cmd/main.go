package main

import (
	"fmt"
	"jameesjohn.com/hngStageTwo/src/config"
	"jameesjohn.com/hngStageTwo/src/database"
	"jameesjohn.com/hngStageTwo/src/routes"
	"log"
	"net/http"
)

func main() {
	//Setup Env
	config.Load()

	// Setup Database Connection
	database.ConnectDatabase()
	database.Migrate()

	// Setup routes
	router := routes.Router()

	log.Println("Http server running on port", config.Environment.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Environment.Port), router))
}
