package main

import (
	"context"
	"log"

	pb "github.com/magodo/shippy-service/user/proto/user"
	"github.com/micro/go-micro"
)

const topic = "user.created"

type subscriber struct{}

func (s *subscriber) Process(ctx context.Context, req *pb.User) error {
	log.Printf("topic %s received\n", topic)
	log.Println(req)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.srv.email"),
	)

	srv.Init()

	if err := micro.RegisterSubscriber(topic, srv.Server(), new(subscriber)); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
