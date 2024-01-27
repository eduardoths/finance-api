package repositories

import (
	"context"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	"github.com/eduardoths/finance-api/pkg/db"
)

type ExpenseRepository struct{}

func NewExpenseRepository() ExpenseRepository {
	return ExpenseRepository{}
}

func (er ExpenseRepository) Create(ctx context.Context, data structs.Expense) (structs.Expense, error) {
	if err := db.GetFromContext(ctx).
		Create(&data).
		Error; err != nil {
		return structs.Expense{}, err
	}
	return data, nil
}
