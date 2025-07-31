package user

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/user"
)

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required,min=3,max=20"`
		Password string `json:"password" validate:"required,min=6,max=50"`
		Email    string `json:"email" validate:"required,email"`
	}
)

func (r RegisterRequest) ToUcModel() ucModel.RegisterRequest {
	return ucModel.RegisterRequest{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
