package main

import (
	"context"
	"log"

	"os"

	"github.com/magodo/shippy-service/consignment/internal"
	pb "github.com/magodo/shippy-service/consignment/proto/consignment"
	vesselPb "github.com/magodo/shippy-service/vessel/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultDBHost = "mongodb://localhost:27017"
)

func main() {

	dbHost := defaultDBHost
	if envDBHost, ok := os.LookupEnv("DB_HOST"); ok {
		dbHost = envDBHost
	}
	client, err := internal.CreateClient(dbHost)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	consignmentCollection := client.Database("shippy").Collection("consignments")
	repo := &internal.Repository{Collection: consignmentCollection}

	srv := micro.NewService(
		micro.Name("shippy.consignment.service"),
	)
	srv.Init()
	vesselClient := vesselPb.NewVesselService("shippy.vessel.service", srv.Client())
	pb.RegisterShippingServiceHandler(srv.Server(), &internal.Handler{Repo: repo, VesselClient: vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
