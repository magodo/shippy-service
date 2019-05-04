package internal

import (
	"context"
	"time"

	pb "github.com/magodo/shippy-service/consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

type Repository struct {
	Collection *mongo.Collection
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) error {
	_, err := repo.Collection.InsertOne(context.Background(), consignment)
	return err
}

func (repo *Repository) GetAll() ([]*pb.Consignment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cur, err := repo.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var consignments []*pb.Consignment
	for cur.Next(ctx) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
