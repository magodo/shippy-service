package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"context"

	pb "github.com/magodo/shippy-service/consignment/proto/consignment"
	"github.com/micro/cli"
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
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "json",
				Usage: "json file of consignment to create",
			},
		),
		micro.Name("shippy.consignment.cli"),
	)

	var file string
	service.Init(
		micro.Action(func(c *cli.Context) {
			file = c.String("json")
		}),
	)

	client := pb.NewShippingService("shippy.consignment.service", service.Client())

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
