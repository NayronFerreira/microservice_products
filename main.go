package main

import (
	"log"

	"github.com/NayronFerreira/microservice_products/configs"
	"github.com/NayronFerreira/microservice_products/internal/infra/database"
	server "github.com/NayronFerreira/microservice_products/internal/infra/web"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	db, err := database.SetupDB(config)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer db.Close()

	log.Println("server is starting...")
	if err := server.NewServer(config, db).SetupServer().ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
