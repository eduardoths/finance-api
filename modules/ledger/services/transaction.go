package services

import (
	"context"
	"errors"
	"time"

	"github.com/eduardoths/finance-api/modules/ledger/dto"
	"github.com/eduardoths/finance-api/modules/ledger/repositories"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
)

type TransactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return TransactionService{
		repo: repo,
	}
}

func (ts TransactionService) New(ctx context.Context, t dto.NewTransaction) (structs.Transaction, error) {

	return ts.repo.New(ctx, structs.Transaction{
		Amount:      t.Amount,
		Date:        t.Date,
		Description: t.Description,
		DebitID:     t.DebitID,
		CreditID:    t.CreditID,
		AssetID:     t.AssetID,
	})
}

func (ts TransactionService) Get(ctx context.Context, id structs.ID) (structs.Transaction, error) {
	return ts.repo.Get(ctx, id)
}

func (ts TransactionService) Undo(ctx context.Context, id structs.ID) (structs.Transaction, error) {
	t, err := ts.Get(ctx, id)
	if err != nil {
		return structs.Transaction{}, err
	}

	if (t.Description) == "ROLLBACK" {
		return structs.Transaction{}, errors.New("Can't rollback after a rollback")
	}

	return ts.New(ctx, dto.NewTransaction{
		Amount:      t.Amount,
		Date:        time.Now(),
		Description: "ROLLBACK",
		AssetID:     t.AssetID,
		DebitID:     t.CreditID,
		CreditID:    t.DebitID,
	})
}

func (ts TransactionService) GetMany(ctx context.Context, filter dto.TransactionFilter) ([]structs.Transaction, error) {
	return ts.repo.GetMany(ctx, filter)
}

func (as TransactionService) GetAll(ctx context.Context) ([]structs.Transaction, error) {
	return as.repo.GetAll(ctx)
}
