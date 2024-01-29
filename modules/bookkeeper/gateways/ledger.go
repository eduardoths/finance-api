package gateways

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger"
	ledgerDTO "github.com/eduardoths/finance-api/modules/ledger/dto"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
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

func (lg LedgerGateway) GetTransaction(ctx context.Context, id ledgerStructs.ID) (ledgerStructs.Transaction, error) {
	return lg.ledger.Transactions.Get(ctx, id)
}

func (lg LedgerGateway) GetAsset(ctx context.Context, id ledgerStructs.ID) (ledgerStructs.Asset, error) {
	// return lg.ledger.Assets.Get(ctx, id)
	return structs.Asset{}, nil
}

func (lg LedgerGateway) Rollback(ctx context.Context, id uint) (ledgerStructs.Transaction, error) {
	return lg.ledger.Transactions.Undo(ctx, id)
}

func (lg LedgerGateway) Commit(ctx context.Context, id uint) error {
	return nil
}
