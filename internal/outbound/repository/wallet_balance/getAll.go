package wallet_balance

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) GetAll(ctx context.Context, orm *gorm.DB, req model.WalletBalance) ([]model.WalletBalance, error) {
	var walletBalances []model.WalletBalance

	orm = orm.WithContext(ctx)

	if err := orm.Where(req).Preload("Currency").Find(&walletBalances).Error; err != nil {
		r.resource.Logger.Errorf("[WalletBalanceRepository] GetAll: %v", err)
		return nil, err
	}

	return walletBalances, nil
}

type IGetAll interface {
	GetAll(ctx context.Context, orm *gorm.DB, req model.WalletBalance) ([]model.WalletBalance, error)
}
