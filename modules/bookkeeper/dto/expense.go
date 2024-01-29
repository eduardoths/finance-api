package dto

import (
	"time"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	ledgerStructs "github.com/eduardoths/finance-api/modules/ledger/structs"
)

type CreateExpense struct {
	Date        time.Time               `json:"date"`
	FromAccount uint                    `json:"from_account"`
	Currency    uint                    `json:"currency"`
	Category    structs.ExpenseCategory `json:"category"`
	Amount      uint                    `json:"amount"`
	Description string                  `json:"description"`
}

type ViewExpense struct {
	ID            uint                      `json:"id"`
	CreatedAt     time.Time                 `json:"created_at"`
	UpdatedAt     time.Time                 `json:"updated_at"`
	Category      string                    `json:"category"`
	Amount        uint                      `json:"amount"`
	Date          time.Time                 `json:"date"`
	Description   string                    `json:"description"`
	TransactionID uint                      `json:"transaction_id"`
	Transaction   ledgerStructs.Transaction `json:"transaction"`
}

func NewViewExpense(exp structs.Expense, tx ledgerStructs.Transaction) ViewExpense {
	return ViewExpense{
		ID:            exp.ID,
		CreatedAt:     exp.CreatedAt,
		UpdatedAt:     exp.UpdatedAt,
		Category:      string(exp.Category),
		Amount:        exp.Amount,
		Date:          exp.Date,
		Description:   exp.Description,
		TransactionID: exp.TransactionID,
		Transaction:   tx,
	}
}
