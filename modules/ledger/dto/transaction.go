package dto

import (
	"time"

	"github.com/eduardoths/finance-api/modules/ledger/structs"
	"gorm.io/gorm"
)

type NewTransaction struct {
	Amount      uint
	Date        time.Time
	Description string
	AssetID     structs.ID
	DebitID     structs.ID
	CreditID    structs.ID
}

type TransactionFilter struct {
	MinDate   *time.Time
	MaxDate   *time.Time
	AccountId *structs.ID
	AssetID   *structs.ID
}

func (t TransactionFilter) PrepareConn(conn *gorm.DB) *gorm.DB {
	if t.MinDate != nil {
		conn = conn.Where("date >= ?", t.MinDate)
	}

	if t.MaxDate != nil {
		conn = conn.Where("date <= ?", t.MaxDate)
	}

	if t.AccountId != nil {
		conn = conn.Where("debit_id = ? OR credit_id = ?", t.AccountId, t.AccountId)
	}
	return conn
}
