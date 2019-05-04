package internal

import (
	"context"

	pb "github.com/magodo/shippy-service/user/proto/user"
)

type Service struct {
	Repo         Repository
	TokenService Authable
}

func (s *Service) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	user, err := s.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	resp.User = user
	return nil
}

func (srv *Service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *Service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := srv.Repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingabc"
	return nil
}

func (srv *Service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *Service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
