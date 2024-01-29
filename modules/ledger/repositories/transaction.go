package repositories

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger/dto"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
	"github.com/eduardoths/finance-api/pkg/db"
	"gorm.io/gorm/clause"
)

type TransactionRepository struct{}

func (ts TransactionRepository) New(ctx context.Context, t structs.Transaction) (structs.Transaction, error) {
	if err := db.GetFromContext(ctx).
		Create(&t).
		Error; err != nil {
		return structs.Transaction{}, err
	}
	return t, nil
}

func (ts TransactionRepository) Get(ctx context.Context, id structs.ID) (structs.Transaction, error) {
	var transaction structs.Transaction
	if err := db.GetFromContext(ctx).
		Preload(clause.Associations).
		Take(&transaction).
		Error; err != nil {
		return structs.Transaction{}, err
	}
	return transaction, nil
}

func (ts TransactionRepository) GetMany(ctx context.Context, filter dto.TransactionFilter) ([]structs.Transaction, error) {
	var transactions []structs.Transaction
	if err := filter.PrepareConn(db.GetFromContext(ctx)).
		Preload(clause.Associations).
		Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (tr TransactionRepository) GetAll(ctx context.Context) ([]structs.Transaction, error) {
	var data []structs.Transaction
	if err := db.GetFromContext(ctx).
		Preload(clause.Associations).
		Find(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}
