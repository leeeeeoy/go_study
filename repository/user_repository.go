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

	// 트랜잭션, 생성 수정 삭제 할 때는 트랜잭션 하는게 좋아요 + Row level Lock
	tx, _ := userRepository.db.Tx(context.TODO())
	tx.Commit()
	tx.Rollback()

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
