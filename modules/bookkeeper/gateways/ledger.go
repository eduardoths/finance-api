package gateways

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger"
	ledgerDTO "github.com/eduardoths/finance-api/modules/ledger/dto"
	ledgerStructs "github.com/eduardoths/finance-api/modules/ledger/structs"
)

type LedgerGateway struct {
	ledger ledger.Ledger
}

func NewLedgerGateway() LedgerGateway {
	return LedgerGateway{
		ledger: ledger.NewLedger(),
	}
}

func (lg LedgerGateway) NewTransaction(ctx context.Context, data ledgerDTO.NewTransaction) (ledgerStructs.Transaction, error) {
	return lg.ledger.Transactions.New(ctx, data)
}

func (lg LedgerGateway) Rollback(ctx context.Context, id uint) (ledgerStructs.Transaction, error) {
	return lg.ledger.Transactions.Undo(ctx, id)
}

func (lg LedgerGateway) Commit(ctx context.Context, id uint) error {
	return nil
}
