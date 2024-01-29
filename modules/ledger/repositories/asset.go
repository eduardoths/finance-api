package repositories

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger/structs"
	"github.com/eduardoths/finance-api/pkg/db"
	"gorm.io/gorm/clause"
)

type AssetRepository struct {
}

func (ar AssetRepository) Create(ctx context.Context, data structs.Asset) (structs.Asset, error) {
	if err := db.GetFromContext(ctx).
		Create(&data).
		Error; err != nil {
		return structs.Asset{}, err
	}
	return data, nil
}

func (ar AssetRepository) Get(ctx context.Context, id structs.ID) (structs.Asset, error) {
	var data structs.Asset
	if err := db.GetFromContext(ctx).
		Preload(clause.Associations).
		Take(&data).
		Error; err != nil {
		return structs.Asset{}, err
	}
	return data, nil
}
