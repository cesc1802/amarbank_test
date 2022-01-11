package domain

import "github.com/google/uuid"

type UserModel struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Dob    string    `json:"dob"`
	Gender string    `json:"gender"`
}

func NewUserModel(name, dob string, gender string) UserModel {
	return UserModel{
		Name:   name,
		Dob:    dob,
		Gender: gender,
	}
}

func (m *UserModel) GenerateID() {
	if m.ID != uuid.Nil {
		return
	}
	m.ID = uuid.New()
}
