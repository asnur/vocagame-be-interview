package transaction

import (
	ucModel "github.com/asnur/vocagame-be-interview/internal/usecase/model/transaction"
)

type (
	TransactionRequest struct {
		TransactionID string `json:"transaction_id"`
	}
)

func (t TransactionRequest) ToUcModel(userId int64) ucModel.TransactionGet {
	return ucModel.TransactionGet{
		UserID: userId,
		TrxID:  t.TransactionID,
	}
}
