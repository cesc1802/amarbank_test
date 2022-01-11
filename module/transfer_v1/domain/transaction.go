package domain

import (
	app_error "amarbank/pkg/apperror"
	"github.com/google/uuid"
)

type TransactionModel struct {
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	Amount     float64   `json:"amount"`
	Date       string    `json:"date"`
}

func NewTransactionModel(senderID, receiverID uuid.UUID, amount float64, date string) TransactionModel {
	return TransactionModel{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     amount,
		Date:       date,
	}
}

func ErrSenderNotFound(err error) *app_error.AppError {
	return app_error.NewCustomError(err, "sender not found", "SENDER_NOT_FOUND")
}

func ErrAccountSenderNotFound(err error) *app_error.AppError {
	return app_error.NewCustomError(err, "account sender not found", "ACCOUNT_NOT_FOUND")
}

func ErrInvalidPin(err error) *app_error.AppError {
	return app_error.NewCustomError(err, "invalid pin", "INVALID_PIN")
}
