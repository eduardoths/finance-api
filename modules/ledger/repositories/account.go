package repositories

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger/structs"
	"github.com/eduardoths/finance-api/pkg/db"
	"gorm.io/gorm/clause"
)

type AccountRepository struct {
}

func (ar AccountRepository) Create(ctx context.Context, data structs.Account) (structs.Account, error) {
	if err := db.GetFromContext(ctx).
		Create(&data).
		Error; err != nil {
		return structs.Account{}, err
	}
	return data, nil
}

func (ar AccountRepository) Get(ctx context.Context, id structs.ID) (structs.Account, error) {
	var data structs.Account
	if err := db.GetFromContext(ctx).
		Preload(clause.Associations).
		Take(&data).
		Error; err != nil {
		return structs.Account{}, err
	}
	return data, nil
}

func (ar AccountRepository) GetAll(ctx context.Context) ([]structs.Account, error) {
	var data []structs.Account
	if err := db.GetFromContext(ctx).
		Preload(clause.Associations).
		Find(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}
