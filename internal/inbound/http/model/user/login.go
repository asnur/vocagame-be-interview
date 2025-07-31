package user

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/user"
)

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

func (u LoginRequest) ToUcModel() ucModel.LoginRequest {
	return ucModel.LoginRequest{
		Username: u.Username,
		Password: u.Password,
	}
}
