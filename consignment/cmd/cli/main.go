package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"context"

	pb "github.com/magodo/shippy-service/consignment/proto/consignment"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
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
			cli.StringFlag{
				Name:  "token",
				Usage: "authorization token",
			},
		),
		micro.Name("shippy.consignment.cli"),
	)

	var (
		file  string
		token string
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			file = c.String("json")
			token = c.String("token")
		}),
	)

	client := pb.NewShippingService("shippy.consignment.service", service.Client())

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), metadata.Metadata{"Token": token})
	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	r, err = client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not get all consignments: %v", err)
	}
	log.Printf("Consignments: %v", r.Consignments)
}
