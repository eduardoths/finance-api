package dto

import (
	"time"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
)

type CreateIncome struct {
	Date        time.Time              `json:"date"`
	ToAccount   structs.ExternalID     `json:"to_account"`
	Currency    structs.ExternalID     `json:"currency"`
	Category    structs.IncomeCategory `json:"category"`
	Amount      uint                   `json:"amount"`
	Description string                 `json:"description"`
}
