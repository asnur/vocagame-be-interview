package user

import (
	"context"
	"strings"

	ucSharedModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/shared"
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/user"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
)

func (u *usecase) Register(ctx context.Context, req ucModel.RegisterRequest) (ucModel.RegisterResponse, error) {
	var (
		orm = u.resource.Postgres.DB
	)

	// Hash the password using shared use case
	hashedPassword, err := u.shared.HashPassword(ctx, ucSharedModel.HashPasswordRequest{Input: req.Password, Key: u.resource.AppConfig.Key})
	if err != nil {
		return ucModel.RegisterResponse{}, err
	}

	req.Password = hashedPassword.Hash

	tx := orm.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		u.resource.Logger.Errorf("[UserUseCase] Register: failed to begin transaction: %v", err)
		return ucModel.RegisterResponse{}, err
	}

	// Create Data User
	_, err = u.Repository.User.Create(ctx, tx, req.ToObUserModel())
	if err != nil {
		u.resource.Logger.Errorf("[UserUseCase] Register: %v", err)
		tx.Rollback()

		if strings.Contains(err.Error(), pkgErr.ErrDuplicate.Error()) {
			return ucModel.RegisterResponse{}, pkgErr.ErrIdentityAlreadyExists
		}

		return ucModel.RegisterResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		u.resource.Logger.Errorf("[UserUseCase] Register: failed to commit transaction: %v", err)
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
