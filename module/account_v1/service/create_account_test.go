package service

import (
	accountDomain "amarbank/module/account_v1/domain"
	"amarbank/module/account_v1/dto"
	userDomain "amarbank/module/user_v1/domain"
	app_error "amarbank/pkg/apperror"
	"amarbank/pkg/valueobject"
	"context"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

type mockCreateAccountStorage struct {
	createAccount func(input *accountDomain.AccountModel) error
}

type mockGetUserStorage struct {
	getUserByID func(uuid uuid.UUID) (*userDomain.UserModel, error)
}

type mockGeneratePIN struct {
	gen func(typ valueobject.GenderType) string
}

func (m *mockCreateAccountStorage) Create(ctx context.Context, input *accountDomain.AccountModel) error {
	if m != nil && m.createAccount != nil {
		return m.createAccount(input)
	}
	return nil
}

func (m *mockGetUserStorage) GetByID(ctx context.Context, id uuid.UUID) (*userDomain.UserModel, error) {
	if m != nil && m.getUserByID != nil {
		return m.getUserByID(id)
	}
	return &userDomain.UserModel{}, nil
}
func (m *mockGeneratePIN) Gen(typ valueobject.GenderType) string {
	if m != nil && m.gen != nil {
		return m.gen(typ)
	}
	return "qwerty11012022qwer"
}

func TestCreateAccountService_CreateAccount(t *testing.T) {
	tests := []struct {
		name          string
		mockPIN       *mockGeneratePIN
		mockUserStore *mockGetUserStorage
		mockAccStore  *mockCreateAccountStorage
		userID        uuid.UUID
		defaultAmount float64
		expectedErr   func(err error) *app_error.AppError
	}{
		{
			name: "test case not found user id",
			mockPIN: &mockGeneratePIN{
				gen: func(typ valueobject.GenderType) string {
					if typ.IsMale() {
						return "qwerty11012022qwer"
					} else if typ.IsMale() {
						return "qwerty51012022qwer"
					}
					return "qwerty11012022qwer"

				},
			},
			mockUserStore: &mockGetUserStorage{
				getUserByID: func(id uuid.UUID) (*userDomain.UserModel, error) {
					return nil, fmt.Errorf("user not found")
				},
			},
			mockAccStore: &mockCreateAccountStorage{
				createAccount: func(input *accountDomain.AccountModel) error {
					return nil
				},
			},
			userID:      uuid.MustParse("f0869d5e-0450-468a-9323-c3421b2a9b67"),
			expectedErr: accountDomain.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCreateAccountService(tt.mockAccStore, tt.mockUserStore, tt.mockPIN)

			input := dto.CreateAccountRequest{
				UserID: tt.userID,
			}
			err := service.CreateAccount(context.TODO(), &input)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if tt.expectedErr(err).Log != errMsg {
				t.Errorf("Unexpected error: %v", errMsg)
			}
		})
	}
}
