package shared

import (
	"context"

	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/shared"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) HashPassword(ctx context.Context, req ucModel.HashPasswordRequest) (ucModel.HashPasswordResponse, error) {
	// Use bcrypt to hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Input), bcrypt.DefaultCost)
	if err != nil {
		return ucModel.HashPasswordResponse{}, err
	}

	return ucModel.HashPasswordResponse{
		Hash: string(hashedPassword),
	}, nil
}

type IHashPassword interface {
	HashPassword(ctx context.Context, req ucModel.HashPasswordRequest) (ucModel.HashPasswordResponse, error)
}
