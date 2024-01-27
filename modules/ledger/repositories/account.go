package repositories

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger/structs"
	"github.com/eduardoths/finance-api/pkg/db"
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
		Take(&data).
		Error; err != nil {
		return structs.Account{}, err
	}
	return data, nil
}
