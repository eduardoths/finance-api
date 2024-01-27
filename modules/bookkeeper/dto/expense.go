package dto

import (
	"time"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
)

type CreateExpense struct {
	Date        time.Time               `json:"date"`
	FromAccount uint                    `json:"from_account"`
	Currency    uint                    `json:"currency"`
	Category    structs.ExpenseCategory `json:"category"`
	Amount      uint                    `json:"amount"`
	Description string                  `json:"description"`
}
