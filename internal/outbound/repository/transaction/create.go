package transaction

import (
	"context"

	"github.com/asnur/vocagame-be-interview/internal/outbound/model"
	"gorm.io/gorm"
)

func (r *repository) Create(ctx context.Context, orm *gorm.DB, transaction model.Transaction) (model.Transaction, error) {
	if err := orm.WithContext(ctx).Create(&transaction).Error; err != nil {
		r.resource.Logger.Errorf("[TransactionRepository] Create: %v", err)
		return model.Transaction{}, err
	}

	return transaction, nil
}

type ICreate interface {
	Create(ctx context.Context, orm *gorm.DB, transaction model.Transaction) (model.Transaction, error)
}
