package service

import (
	accountDomain "amarbank/module/account_v1/domain"
	"amarbank/module/transfer_v1/domain"
	"amarbank/module/transfer_v1/dto"
	userDomain "amarbank/module/user_v1/domain"
	app_error "amarbank/pkg/apperror"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockGetAccountStorage struct {
	getAccByUserIDFn func(userUuid uuid.UUID) (*accountDomain.AccountModel, error)
}

type mockGetUserStorage struct {
	mock.Mock
	getUserByID func(uuid uuid.UUID) (*userDomain.UserModel, error)
}

type mockCreateTxnStorage struct {
	mock.Mock
	createTxn func(input *domain.TransactionModel) error
}

func (ma *mockGetAccountStorage) GetByUserID(ctx context.Context, userUuid uuid.UUID) (*accountDomain.AccountModel, error) {
	if ma != nil && ma.getAccByUserIDFn != nil {
		return ma.getAccByUserIDFn(userUuid)
	}
	return &accountDomain.AccountModel{}, nil
}

func (mu *mockGetUserStorage) GetByID(ctx context.Context, uuid uuid.UUID) (*userDomain.UserModel, error) {
	if mu != nil && mu.getUserByID != nil {
		return mu.getUserByID(uuid)
	}
	return &userDomain.UserModel{}, nil
}
func (mtxn *mockCreateTxnStorage) Create(ctx context.Context, input *domain.TransactionModel) error {
	if mtxn != nil && mtxn.createTxn != nil {
		return mtxn.createTxn(input)
	}
	return nil
}

func TestCreateTransfer(t *testing.T) {

	tests := []struct {
		testName           string
		mockAccStore       *mockGetAccountStorage
		mockUserStore      *mockGetUserStorage
		mockTxnStore       *mockCreateTxnStorage
		SenderID           string
		ReceiverID         string
		KTP                string
		LoanAmount         float64
		LoanPeriodInMonths string
		LoanPurpose        string
		expectedError      func(err error) *app_error.AppError
	}{
		{
			testName: "test case not found senderID",
			mockAccStore: &mockGetAccountStorage{
				getAccByUserIDFn: func(userUuid uuid.UUID) (*accountDomain.AccountModel, error) {
					return nil, nil
				},
			},
			mockUserStore: &mockGetUserStorage{
				getUserByID: func(uuid uuid.UUID) (*userDomain.UserModel, error) {
					return nil, fmt.Errorf("not found user")
				},
			},

			mockTxnStore: &mockCreateTxnStorage{
				createTxn: func(input *domain.TransactionModel) error {
					return nil
				},
			},

			SenderID:           "f0869d5e-0450-468a-9323-c3421b2a9b3d",
			ReceiverID:         "f0869d5e-0450-468a-9323-c3421b2a9b3f",
			KTP:                "5dPumk100122qPMM",
			LoanAmount:         8888,
			LoanPeriodInMonths: "2",
			LoanPurpose:        "vacation",
			expectedError:      domain.ErrSenderNotFound,
		},

		{
			testName: "test case invalid pin",
			mockAccStore: &mockGetAccountStorage{
				getAccByUserIDFn: func(userUuid uuid.UUID) (*accountDomain.AccountModel, error) {
					return &accountDomain.AccountModel{
						ID:     uuid.MustParse("08c0695e-3132-4608-8778-a40d60dd9667"),
						UserID: uuid.MustParse("8f0a687d-6eb6-415a-bd93-618c241a3004"),
						Amount: 50000,
						Pin:    "5dPumk100122qPMM",
					}, nil
				},
			},
			mockUserStore: &mockGetUserStorage{
				getUserByID: func(id uuid.UUID) (*userDomain.UserModel, error) {
					return &userDomain.UserModel{
						ID:     uuid.MustParse("f0869d5e-0450-468a-9323-c3421b2a9b3f"),
						Name:   "alex nguyen",
						Dob:    "18-02-1993",
						Gender: "male",
					}, nil
				},
			},

			mockTxnStore: &mockCreateTxnStorage{
				createTxn: func(input *domain.TransactionModel) error {
					return nil
				},
			},

			SenderID:           "f0869d5e-0450-468a-9323-c3421b2a9b3f",
			ReceiverID:         "f0869d5e-0450-468a-9323-c3421b2a9b3f",
			KTP:                "5dPumk100122qPMM",
			LoanAmount:         8888,
			LoanPeriodInMonths: "2",
			LoanPurpose:        "vacation",
			expectedError:      domain.ErrInvalidPin,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			service := NewCreateTxnService(tt.mockTxnStore, tt.mockAccStore, tt.mockUserStore)

			input := dto.CreateTransferRequest{
				SenderID:           uuid.MustParse(tt.SenderID),
				ReceiverID:         uuid.MustParse(tt.ReceiverID),
				KTP:                tt.KTP,
				LoanAmount:         tt.LoanAmount,
				LoanPeriodInMonths: tt.LoanPeriodInMonths,
				LoanPurpose:        tt.LoanPurpose,
			}
			err := service.CreateTxn(context.TODO(), &input)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if tt.expectedError(err).Log != errMsg {
				t.Errorf("Unexpected error: %v", errMsg)
			}

		})
	}

}
