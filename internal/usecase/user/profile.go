package user

import (
	"context"
	"errors"

	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/user"
	pkgErr "github.com/asnur/vocagame-be-interview/pkg/errors"
	"gorm.io/gorm"
)

func (u *usecase) Profile(ctx context.Context, req ucModel.ProfileRequest) (ucModel.ProfileResponse, error) {
	var (
		orm      = u.resource.Postgres.DB
		response = ucModel.ProfileResponse{}
	)
	// Get user data from repository
	user, err := u.Repository.User.Get(ctx, orm, req.ToObUserModel())
	if err != nil {
		u.resource.Logger.Errorf("[UserUseCase] Profile: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ucModel.ProfileResponse{}, pkgErr.ErrUserNotFound
		}

		return ucModel.ProfileResponse{}, err
	}

	// prepare response
	response.ID = user.ID
	response.Username = user.Username
	response.Email = user.Email
	for _, wallet := range user.Wallets {
		response.Wallets = append(response.Wallets, ucModel.ProfileWallet{
			ID:      wallet.ID,
			Name:    wallet.Name,
			Balance: 0,
		})
	}

	return response, nil
}

type IProfile interface {
	Profile(ctx context.Context, req ucModel.ProfileRequest) (ucModel.ProfileResponse, error)
}
