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
	return nil, nil
	// tx, err := userRepository.db.Tx(context.Background())

	// if err != nil {
	// return nil, err
	// }

	// res, err := tx.User.Create().
	// 	SetEmail(userRequest.Email).
	// 	SetName(userRequest.Name).
	// 	SetPassword(userRequest.Password).
	// 	Save(context.Background())

	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	// result := &dto.UserResponse{
	// 	ID:        res.ID,
	// 	Email:     res.Email,
	// 	Name:      res.Name,
	// 	CreatedAt: res.CreatedAt,
	// }

	// if err := tx.Commit(); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	// return result, nil

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
