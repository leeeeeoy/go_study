package repository

import (
	"context"

	"github.com/leeeeeoy/go_study/dto"
	"github.com/leeeeeoy/go_study/ent"
	"github.com/leeeeeoy/go_study/ent/user"
)

type UserRepository interface {
	CreateUser(userRequest *dto.UserRequest) (*dto.UserResponse, error)
	GetUserByEmail(email string) (*dto.UserResponse, error)
	SignIn(email, password string) (string, error)
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

	tx, err := userRepository.db.Tx(context.Background())

	if err != nil {
		return nil, err
	}

	res, err := userRepository.db.User.Create().
		SetEmail(userRequest.Email).
		SetName(userRequest.Name).
		SetPassword(userRequest.Password).
		Save(context.Background())

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	result := &dto.UserResponse{
		ID:        res.ID,
		Email:     res.Email,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return result, nil

}

func (userRepository *userRepositoryImpl) GetUserByEmail(email string) (*dto.UserResponse, error) {
	res, err := userRepository.db.User.Query().Where(user.Email(email)).Only(context.Background())

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

func (userRepository *userRepositoryImpl) SignIn(email, password string) (string, error) {
	res, err := userRepository.db.User.Query().Where(user.Email(email), user.Password(password)).Only(context.Background())

	if err != nil {
		return "", err
	}

	if res.Password != password {
		return "", err
	}

	return res.Name, nil
}
