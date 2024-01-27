package structs

import (
	"time"
)

type Transaction struct {
	Model
	Amount      uint
	Date        time.Time
	Description string

	DebitID ID
	Debit   Account

	CreditID ID
	Credit   Account

	Asset   Asset
	AssetID ID
}

type TransactionType uint

const (
	DEBIT TransactionType = iota
	CREDIT
)
