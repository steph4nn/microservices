package main

import (

	"github.com/steph4nn/microservices/shipping/config"
	"github.com/steph4nn/microservices/shipping/internal/adapters/db"
	"github.com/steph4nn/microservices/shipping/internal/adapters/grpc"
	"github.com/steph4nn/microservices/shipping/internal/application/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
