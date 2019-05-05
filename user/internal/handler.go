package internal

import (
	"context"
	"errors"
	"log"

	pb "github.com/magodo/shippy-service/user/proto/user"
	"golang.org/x/crypto/bcrypt"
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
	log.Printf("Logging in with: email: %s, password: %s", req.Email, req.Password)
	user, err := srv.Repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	log.Println(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	token, err := srv.TokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *Service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashed)
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *Service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := srv.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true
	return nil
}
