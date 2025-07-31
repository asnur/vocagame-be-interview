package wallet

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repository) Create(ctx context.Context, orm *gorm.DB, wallet model.Wallets) (model.Wallets, error) {
	if err := orm.WithContext(ctx).Preload(clause.Associations).Create(&wallet).Error; err != nil {
		r.resource.Logger.Errorf("[WalletRepository] Create: %v", err)
		return model.Wallets{}, err
	}

	return wallet, nil
}

type ICreate interface {
	Create(ctx context.Context, orm *gorm.DB, wallet model.Wallets) (model.Wallets, error)
}
