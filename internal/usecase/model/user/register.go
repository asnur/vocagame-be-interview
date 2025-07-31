package user

import (
	obModel "github.com/asnur/vocagame-be-interview/internal/outbound/model"
)

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	RegisterResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)

func (u RegisterRequest) ToObUserModel() obModel.Users {
	return obModel.Users{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
}
