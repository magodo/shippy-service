package main

import (
	"log"

	"github.com/magodo/shippy-service/user/internal"
	pb "github.com/magodo/shippy-service/user/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	db, err := internal.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&pb.User{})

	repo := &internal.UserRepository{DB: db}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &internal.Service{Repo: repo})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
