package service

import (
	"amarbank/module/user_v1/domain"
	"amarbank/module/user_v1/dto"
	"context"
	"github.com/google/uuid"
)

type GetUserFileStorage interface {
	GetByID(ctx context.Context, uuid uuid.UUID) (*domain.UserModel, error)
}

type getUserByIDService struct {
	store GetUserFileStorage
}

func NewGetUserByIDService(store GetUserFileStorage) *getUserByIDService {
	return &getUserByIDService{
		store: store,
	}
}

func (s *getUserByIDService) GetUserByID(ctx context.Context, userUUID uuid.UUID) (*dto.GetUserByUserIDResponse, error) {
	userModel, err := s.store.GetByID(ctx, userUUID)

	if err != nil {
		return nil, err
	}

	userResponse := dto.NewGetUserByUserIDResponse(userModel.Name, userModel.Dob, userModel.Gender)
	return &userResponse, nil
}
