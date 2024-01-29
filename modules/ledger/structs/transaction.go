package structs

import (
	"time"
)

type Transaction struct {
	Model
	Amount      uint      `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	DebitID     ID        `json:"debit_id"`
	Debit       Account   `json:"debit"`
	CreditID    ID        `json:"credit_id"`
	Credit      Account   `json:"credit"`
	Asset       Asset     `json:"asset"`
	AssetID     ID        `json:"asset_id"`
}
