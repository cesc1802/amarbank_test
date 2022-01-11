package dto

import (
	"github.com/google/uuid"
	"testing"
)

func TestCreateTransferRequest(t *testing.T) {
	tests := []struct {
		name               string
		SenderID           uuid.UUID `json:"sender_id" binding:"required"`
		ReceiverID         uuid.UUID `json:"receiver_id" binding:"required"`
		KTP                string    `json:"ktp" binding:"omitempty"`
		LoanAmount         float64   `json:"loan_amount" binding:"required,min=1000,max=10000"`
		LoanPeriodInMonths string    `json:"loan_period_in_months" binding:"required"`
		LoanPurpose        string    `json:"loan_purpose" binding:"required,oneof='vacation' 'renovation' 'electronics' 'wedding' 'rent' 'car' 'investment'"`
	}{
		{
			name: "test case required field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
