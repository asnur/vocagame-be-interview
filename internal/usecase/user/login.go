package user

import (
	"context"
	"fmt"

	ucSharedModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/shared"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/user"
)

func (u *usecase) Login(ctx context.Context, req ucModel.LoginRequest) (ucModel.LoginResponse, error) {
	// Get Detail User
	user, err := u.Repository.User.Get(ctx, u.resource.Postgres.DB, req.ToObModel())
	if err != nil {
		u.resource.Logger.Errorf("[UserUseCase] Login: %v", err)
		return ucModel.LoginResponse{}, fmt.Errorf("user not found")
	}

	// Validate user credentials
	validatyPassword, err := u.shared.CheckPassword(ctx, ucSharedModel.CheckPasswordRequest{HashedPassword: user.Password, Password: req.Password})
	if err != nil {
		return ucModel.LoginResponse{}, err
	}

	if !validatyPassword.IsValid {
		u.resource.Logger.Errorf("[UserUseCase] Login: Invalid password")
		return ucModel.LoginResponse{}, fmt.Errorf("invalid password")
	}

	// Generate JWT Token
	token, err := u.shared.AuthToken(ctx, ucSharedModel.AuthTokenRequest{
		UserId: user.ID,
	})
	if err != nil {
		u.resource.Logger.Errorf("[UserUseCase] Login: %v", err)
		return ucModel.LoginResponse{}, fmt.Errorf("failed to generate token")
	}

	return ucModel.LoginResponse{
		AccessToken:  token.Token.AccessToken,
		RefreshToken: token.Token.RefreshToken,
	}, nil
}

type ILogin interface {
	Login(ctx context.Context, req ucModel.LoginRequest) (ucModel.LoginResponse, error)
}
