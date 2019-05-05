package internal

import (
	"github.com/jinzhu/gorm"
	pb "github.com/magodo/shippy-service/user/proto/user"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{}
	if err := repo.DB.Where("email = ?", email).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
