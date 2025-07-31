package wallet

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/wallet"
)

type (
	CreateWalletRequest struct {
		Name string `json:"name" validate:"required,min=3,max=50"`
	}
)

func (w CreateWalletRequest) ToUcModel(userID int64) ucModel.CreateWalletRequest {
	return ucModel.CreateWalletRequest{
		Name:   w.Name,
		UserID: userID,
	}
}
