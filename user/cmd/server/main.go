package main

import (
	"log"

	"github.com/magodo/shippy-service/user/internal"
	pb "github.com/magodo/shippy-service/user/proto/user"
	"github.com/micro/go-micro"
)

const (
	topic = "user.created"
)

func main() {
	// prepare repo
	db, err := internal.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&pb.User{})

	repo := &internal.UserRepository{DB: db}

	// prepare service
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	publisher := micro.NewPublisher(topic, srv.Client())
	pb.RegisterUserServiceHandler(
		srv.Server(),
		&internal.Service{
			Repo:         repo,
			TokenService: &internal.TokenService{},
			Publisher:    publisher,
		},
	)

	// run
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
