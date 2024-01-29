package dto

import (
	"time"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	ledgerStructs "github.com/eduardoths/finance-api/modules/ledger/structs"
)

type CreateIncome struct {
	Date        time.Time              `json:"date"`
	ToAccount   structs.ExternalID     `json:"to_account"`
	Currency    structs.ExternalID     `json:"currency"`
	Category    structs.IncomeCategory `json:"category"`
	Amount      uint                   `json:"amount"`
	Description string                 `json:"description"`
}

type ViewIncome struct {
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

func NewViewIncome(inc structs.Income, tx ledgerStructs.Transaction) ViewIncome {
	return ViewIncome{
		ID:            inc.ID,
		CreatedAt:     inc.CreatedAt,
		UpdatedAt:     inc.UpdatedAt,
		Category:      string(inc.Category),
		Amount:        inc.Amount,
		Date:          inc.Date,
		Description:   inc.Description,
		TransactionID: inc.TransactionID,
		Transaction:   tx,
	}
}
