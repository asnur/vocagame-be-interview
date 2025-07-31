package wallet_balance

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repository) Get(ctx context.Context, orm *gorm.DB, req model.WalletBalance) (model.WalletBalance, error) {
	var walletBalance model.WalletBalance

	orm = orm.WithContext(ctx)

	if req.Locking {
		orm = orm.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	if err := orm.Where(req).First(&walletBalance).Error; err != nil {
		r.resource.Logger.Errorf("[WalletBalanceRepository] Get: %v", err)
		return model.WalletBalance{}, err
	}

	return walletBalance, nil
}

type IGet interface {
	Get(ctx context.Context, orm *gorm.DB, req model.WalletBalance) (model.WalletBalance, error)
}
