package repositories

import (
	"context"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	"github.com/eduardoths/finance-api/pkg/db"
)

type IncomeRepository struct{}

func NewIncomeRepository() IncomeRepository {
	return IncomeRepository{}
}

func (ir IncomeRepository) Create(ctx context.Context, data structs.Income) (structs.Income, error) {
	if err := db.GetFromContext(ctx).
		Create(&data).
		Error; err != nil {
		return structs.Income{}, err
	}
	return data, nil
}
