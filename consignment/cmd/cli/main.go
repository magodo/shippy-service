package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/magodo/shippy-service/consignment/proto/consignment"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := grpc.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()
	client := pb.NewShippingService("shippy.consignment.service", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	r, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not get all consignments: %v", err)
	}
	log.Printf("Consignments: %v", r.Consignments)
}
