package main

import (
	"github.com/Manikiran744/microservices/payment/config"
	"github.com/Manikiran744/microservices/payment/internal/adapters/db"
	"github.com/Manikiran744/microservices/payment/internal/adapters/grpc"
	"github.com/Manikiran744/microservices/payment/internal/application/core/api"
	log "github.com/sirupsen/logrus"
)


func main() {

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}