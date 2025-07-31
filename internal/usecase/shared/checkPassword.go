package shared

import (
	"context"

	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/shared"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) CheckPassword(ctx context.Context, req ucModel.CheckPasswordRequest) (ucModel.CheckPasswordResponse, error) {
	// Use bcrypt to compare the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(req.HashedPassword), []byte(req.Password))
	if err != nil {
		return ucModel.CheckPasswordResponse{}, err
	}

	return ucModel.CheckPasswordResponse{
		IsValid: true,
	}, nil
}

type ICheckPassword interface {
	CheckPassword(ctx context.Context, req ucModel.CheckPasswordRequest) (ucModel.CheckPasswordResponse, error)
}
