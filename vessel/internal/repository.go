package internal

import (
	"context"
	"time"

	pb "github.com/magodo/shippy-service/vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
}

type VesselRepository struct {
	Collection *mongo.Collection
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.D{
		{
			Key: "capacity",
			Value: bson.D{
				{
					Key:   "$gt",
					Value: spec.Capacity,
				},
			},
		},
	}
	var vessel pb.Vessel
	repo.Collection.FindOne(context.Background(), filter).Decode(&vessel)
	return &vessel, nil
}
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := repo.Collection.InsertOne(ctx, vessel)
	return err
}
