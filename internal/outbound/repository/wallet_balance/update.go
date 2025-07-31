package wallet_balance

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) Update(ctx context.Context, orm *gorm.DB, walletbalance model.WalletBalance) (model.WalletBalance, error) {
	if err := orm.WithContext(ctx).Where(model.WalletBalance{
		WalletID:   walletbalance.WalletID,
		CurrencyID: walletbalance.CurrencyID,
	}).Updates(&walletbalance).Error; err != nil {
		r.resource.Logger.Errorf("[WalletBalanceRepository] Update: %v", err)
		return model.WalletBalance{}, err
	}

	return walletbalance, nil
}

type IUpdate interface {
	Update(ctx context.Context, orm *gorm.DB, walletbalance model.WalletBalance) (model.WalletBalance, error)
}
