package repository

import (
	"context"
	"database/sql"

	"github.com/leeeeeoy/go_study/dto"
	"github.com/leeeeeoy/go_study/tutorial"
)

type UserRepository interface {
	CreateUser(userRequest *dto.UserRequest) (*dto.UserResponse, error)
	SignIn(email, password string) (string, error)
	GetUsers() ([]*dto.UserResponse, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(client *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: client,
	}
}

func (userRepository *userRepositoryImpl) CreateUser(userRequest *dto.UserRequest) (*dto.UserResponse, error) {
	tx, err := userRepository.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	queries := tutorial.New(userRepository.db)
	qtx := queries.WithTx(tx)

	res, err := qtx.CreateUser(context.TODO(), &tutorial.CreateUserParams{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Username: userRequest.Name,
	})

	if err != nil {
		return nil, err
	}

	result := &dto.UserResponse{
		ID:        res.ID,
		Email:     res.Email,
		Name:      res.Username,
		CreatedAt: res.CreatedAt,
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil

}

func (userRepository *userRepositoryImpl) GetUsers() ([]*dto.UserResponse, error) {
	queries := tutorial.New(userRepository.db)

	res, err := queries.GetUsers(context.TODO())

	if err != nil {
		return nil, err
	}

	var results []*dto.UserResponse

	for _, data := range res {
		results = append(
			results, &dto.UserResponse{
				ID:        data.ID,
				Email:     data.Email,
				Name:      data.Username,
				CreatedAt: data.CreatedAt,
			})
	}

	return results, nil
}

func (userRepository *userRepositoryImpl) SignIn(email, password string) (string, error) {
	queries := tutorial.New(userRepository.db)

	res, err := queries.SignInByEmail(
		context.TODO(),
		&tutorial.SignInByEmailParams{
			Email:    email,
			Password: password,
		})

	if err != nil {
		return "", err
	}

	if res.Password != password {
		return "", err
	}

	return res.Username, nil
}
