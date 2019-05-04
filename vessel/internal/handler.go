package internal

import (
	"context"

	pb "github.com/magodo/shippy-service/vessel/proto/vessel"
)

type Handler struct {
	Repo Repository
}

func (s *Handler) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	vessel, err := s.Repo.FindAvailable(req)
	if err != nil {
		return err
	}

	resp.Vessel = vessel
	return nil
}

func (s *Handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	if err := s.Repo.Create(req); err != nil {
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}
