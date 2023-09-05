package repository

import (
	"context"
	"errors"

	"github.com/leeeeeoy/go_study/dto"
	"github.com/leeeeeoy/go_study/ent"
	"github.com/leeeeeoy/go_study/ent/user"
)

type UserRepository interface {
	CreateUser(userRequest *dto.UserRequest) (*dto.UserResponse, error)
	GetUserByEmail(email *string) (*dto.UserResponse, error)
}

type userRepositoryImpl struct {
	db *ent.Client
}

func NewUserRepository(client *ent.Client) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: client,
	}
}

func (userRepository *userRepositoryImpl) CreateUser(userRequest *dto.UserRequest) (*dto.UserResponse, error) {
	res, err := userRepository.db.User.Create().
		SetEmail(userRequest.Email).
		SetName(userRequest.Name).
		SetPassword(userRequest.Password).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	result := &dto.UserResponse{
		ID:        res.ID,
		Email:     res.Email,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
	}

	return result, nil

}

func (userRepository *userRepositoryImpl) GetUserByEmail(email *string) (*dto.UserResponse, error) {
	res, err := userRepository.db.User.Query().Where(user.Email(*email)).All(context.Background())

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		err = errors.New("유저가 존재하지 않습니다")
		return nil, err
	}

	user := res[0]

	result := &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return result, nil
}
