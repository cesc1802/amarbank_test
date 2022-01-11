package dto

type GetUserByUserIDResponse struct {
	Name   string `json:"name"`
	Dob    string `json:"dob"`
	Gender string `json:"gender"`
}

func NewGetUserByUserIDResponse(name, dob string, gender string) GetUserByUserIDResponse {
	return GetUserByUserIDResponse{
		Name:   name,
		Dob:    dob,
		Gender: gender,
	}
}
