package wallet

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) Get(ctx context.Context, orm *gorm.DB, req model.Wallets) (model.Wallets, error) {
	var wallet model.Wallets
	if err := orm.WithContext(ctx).Preload("WalletBalance").Where(req).First(&wallet).Error; err != nil {
		r.resource.Logger.Errorf("[WalletRepository] Get: %v", err)
		return model.Wallets{}, err
	}

	return wallet, nil
}

type IGet interface {
	Get(ctx context.Context, orm *gorm.DB, req model.Wallets) (model.Wallets, error)
}
