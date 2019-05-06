package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/magodo/shippy-service/vessel/internal"
	pb "github.com/magodo/shippy-service/vessel/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultDBHost = "mongodb://localhost:27017"
)

func createDummyData(repo internal.Repository) {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {
	dbHost := defaultDBHost
	if envDBHost, ok := os.LookupEnv("DB_HOST"); ok {
		dbHost = envDBHost
	}
	if strings.HasPrefix(dbHost, "mongodb://") {
		fmt.Println("yes")
	}

	client, err := internal.CreateClient(dbHost)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database("shippy").Collection("vessel")

	repo := &internal.VesselRepository{Collection: collection}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("shippy.srv.vessel"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &internal.Handler{Repo: repo})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
