package main

import (
	"context"
	"errors"
	"log"

	"os"

	"github.com/magodo/shippy-service/consignment/internal"
	pb "github.com/magodo/shippy-service/consignment/proto/consignment"
	userPb "github.com/magodo/shippy-service/user/proto/user"
	vesselPb "github.com/magodo/shippy-service/vessel/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)

const (
	defaultDBHost = "mongodb://localhost:27017"
)

func main() {

	// connect to datastore and prepare repo
	dbHost := defaultDBHost
	if envDBHost, ok := os.LookupEnv("DB_HOST"); ok {
		dbHost = envDBHost
	}
	mongoClient, err := internal.CreateClient(dbHost)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.Background())
	consignmentCollection := mongoClient.Database("shippy").Collection("consignments")
	repo := &internal.Repository{Collection: consignmentCollection}

	// prepare vessel client
	vesselClient := vesselPb.NewVesselService("shippy.srv.vessel", client.DefaultClient)

	// prepare service
	srv := micro.NewService(
		micro.Name("shippy.srv.consignment"),
		micro.WrapHandler(AuthWrapper),
	)
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &internal.Handler{Repo: repo, VesselClient: vesselClient})

	// run
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// AuthWrapper is a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our consignment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("ENABLE_AUTH") == "false" {
			return fn(ctx, req, resp)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// authenticatet the token
		authClient := userPb.NewUserService("shippy.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		if err != nil {
			log.Println(err)
			return err
		}

		// call the actual request
		return fn(ctx, req, resp)
	}
}
