package service

import (
	accountDomain "amarbank/module/account_v1/domain"
	"amarbank/module/transfer_v1/domain"
	"amarbank/module/transfer_v1/dto"
	userModel "amarbank/module/user_v1/domain"
	app_error "amarbank/pkg/apperror"
	"context"
	"github.com/google/uuid"
	"time"
)

type GetAccountStorage interface {
	GetByUserID(ctx context.Context, userUuid uuid.UUID) (*accountDomain.AccountModel, error)
}

type GetUserStorage interface {
	GetByID(ctx context.Context, uuid uuid.UUID) (*userModel.UserModel, error)
}

type CreateTxnStorage interface {
	Create(ctx context.Context, input *domain.TransactionModel) error
}

type createTxnService struct {
	store        CreateTxnStorage
	accountStore GetAccountStorage
	userStore    GetUserStorage
}

func NewCreateTxnService(store CreateTxnStorage,
	accStore GetAccountStorage,
	userStore GetUserStorage) *createTxnService {
	return &createTxnService{
		store:        store,
		accountStore: accStore,
		userStore:    userStore,
	}
}

func (s *createTxnService) CreateTxn(ctx context.Context, input *dto.CreateTransferRequest) error {
	senderModel, err := s.userStore.GetByID(ctx, input.SenderID)
	if err != nil {
		return domain.ErrSenderNotFound(err)
	}

	accSenderModel, err := s.accountStore.GetByUserID(ctx, senderModel.ID)
	if err != nil {
		return domain.ErrAccountSenderNotFound(err)
	}

	//TODO: check pin
	if accSenderModel.CheckPIN(input.KTP) {
		return domain.ErrInvalidPin(err)
	}

	//TODO: check balance
	if accSenderModel.CheckAmount(input.LoanAmount) {
		return app_error.NewCustomError(err, "amount not enough", "AMOUNT_NOT_ENOUGH")
	}

	receiverModel, err := s.userStore.GetByID(ctx, input.ReceiverID)
	if err != nil {
		return app_error.NewCustomError(err, "receiver not found", "RECEIVER_NOT_FOUND")
	}

	txnModel := domain.NewTransactionModel(senderModel.ID, receiverModel.ID, input.LoanAmount, time.Now().String())
	if err := s.store.Create(ctx, &txnModel); err != nil {
		return app_error.NewCustomError(err, "cannot transfer", "ERR_CANNOT_CREATE_TRANSACTION")
	}

	return nil
}
