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

func (er ExpenseRepository) GetAll(ctx context.Context) ([]structs.Expense, error) {
	var data []structs.Expense
	if err := db.GetFromContext(ctx).
		Find(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}
