package service

import (
	"amarbank/module/user_v1/domain"
	"amarbank/module/user_v1/dto"
	"context"
)

type createUserStorage interface {
	Create(ctx context.Context, user *domain.UserModel) error
}

type createUserService struct {
	store createUserStorage
}

func NewCreateUserService(store createUserStorage) *createUserService {
	return &createUserService{
		store: store,
	}
}

func (s *createUserService) CreateUser(ctx context.Context, user *dto.CreateUserRequest) error {
	userModel := domain.NewUserModel(user.Name, user.Dob, user.Gender)

	userModel.GenerateID()
	if err := s.store.Create(ctx, &userModel); err != nil {
		return err
	}
	return nil
}
