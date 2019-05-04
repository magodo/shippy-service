package main

import (
	"context"
	"log"
	"os"

	pb "github.com/magodo/shippy-service/user/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	client := pb.NewUserService("go.micro.srv.user", service.Client())

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")

			r, err := client.Create(context.Background(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Created: %s", r.User.Id)

			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatal(err)
			}

			for _, v := range getAll.Users {
				log.Println(v)
			}

			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
