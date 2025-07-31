package user

import (
	"context"

	ucSharedModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/shared"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/user"
)

func (u *usecase) Register(ctx context.Context, req ucModel.RegisterRequest) (ucModel.RegisterResponse, error) {
	// Hash the password using shared use case
	hashedPassword, err := u.shared.HashPassword(ctx, ucSharedModel.HashPasswordRequest{Input: req.Password, Key: u.resource.AppConfig.Key})
	if err != nil {
		return ucModel.RegisterResponse{}, err
	}

	req.Password = hashedPassword.Hash

	// Input Data User
	if _, err := u.Repository.User.Create(ctx, u.resource.Postgres.DB, req.ToObModel()); err != nil {
		u.resource.Logger.Errorf("[UserUseCase] Register: %v", err)
		return ucModel.RegisterResponse{}, err
	}

	return ucModel.RegisterResponse{
		Username: req.Username,
		Email:    req.Email,
	}, nil
}

type IRegister interface {
	Register(ctx context.Context, req ucModel.RegisterRequest) (ucModel.RegisterResponse, error)
}
