package dto

type CreateUserRequest struct {
	Name   string `json:"name" binding:"required,space-between"`
	Dob    string `json:"dob" binding:"required"`
	Gender string `json:"gender" binding:"required,oneof='male' 'female' 'MALE' 'FEMALE'"`
}
