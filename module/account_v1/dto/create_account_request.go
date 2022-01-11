package dto

import "github.com/google/uuid"

type CreateAccountRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}
