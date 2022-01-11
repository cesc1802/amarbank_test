package service

import (
	accountDomain "amarbank/module/account_v1/domain"
	"amarbank/module/account_v1/dto"
	userDomain "amarbank/module/user_v1/domain"
	"amarbank/pkg/valueobject"
	"context"
	"github.com/google/uuid"
)

type CreateAccountStorage interface {
	Create(ctx context.Context, input *accountDomain.AccountModel) error
}

type GetUserStorage interface {
	GetByID(ctx context.Context, uuid uuid.UUID) (*userDomain.UserModel, error)
}

type GeneratePIN interface {
	Gen(typ valueobject.GenderType) string
}

type createAccountService struct {
	store     CreateAccountStorage
	userStore GetUserStorage
	pin       GeneratePIN
}

func NewCreateAccountService(store CreateAccountStorage,
	userStore GetUserStorage,
	pin GeneratePIN) *createAccountService {
	return &createAccountService{
		store:     store,
		userStore: userStore,
		pin:       pin,
	}
}

func (s *createAccountService) CreateAccount(ctx context.Context, input *dto.CreateAccountRequest) error {
	userModel, err := s.userStore.GetByID(ctx, input.UserID)
	if err != nil {
		return accountDomain.ErrUserNotFound(err)
	}

	//TODO: we can check some business for user here. Something like user has been band or user has been locked
	// if userModel.Status != Active { return fmt.Error("user has been blocked") }

	accountModel := accountDomain.NewAccountModel(input.UserID)
	accountModel.GenerateID()
	accountModel.SetPin(s.pin.Gen(valueobject.GenderFromString(userModel.Gender)))

	//TODO: always set default amount for any account when created successfully
	accountModel.SetAmount(accountDomain.DefaultAmount)

	if err := s.store.Create(ctx, &accountModel); err != nil {
		return accountDomain.ErrCannotCreateAccount(err)
	}

	return nil
}
