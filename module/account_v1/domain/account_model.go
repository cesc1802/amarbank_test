package domain

import (
	app_error "amarbank/pkg/apperror"
	"github.com/google/uuid"
)

const DefaultAmount = 50000

type AccountModel struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Amount float64   `json:"amount"`
	Pin    string    `json:"pin"`
}

func NewAccountModel(userId uuid.UUID) AccountModel {
	return AccountModel{
		UserID: userId,
	}
}

func (m *AccountModel) GenerateID() {
	if m.ID != uuid.Nil {
		return
	}
	m.ID = uuid.New()
}

func (m *AccountModel) SetPin(pin string) {
	m.Pin = pin
}

func (m *AccountModel) SetAmount(amt float64) {
	m.Amount = amt
}

func (m AccountModel) CheckAmount(transferAmt float64) bool {
	return m.Amount > transferAmt
}

func (m AccountModel) CheckPIN(pin string) bool {
	return m.Pin == pin
}

func ErrUserNotFound(err error) *app_error.AppError {
	return app_error.NewCustomError(err, "user not found", "ERR_USER_NOT_FOUND")
}

func ErrCannotCreateAccount(err error) *app_error.AppError {
	return app_error.NewCustomError(err, "cannot create account", "ERR_CANNOT_CREATE_ACCOUNT")
}
