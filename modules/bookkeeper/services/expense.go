package services

import (
	"context"

	"github.com/eduardoths/finance-api/modules/bookkeeper/domain"
	"github.com/eduardoths/finance-api/modules/bookkeeper/dto"
	"github.com/eduardoths/finance-api/modules/bookkeeper/gateways"
	"github.com/eduardoths/finance-api/modules/bookkeeper/repositories"
	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	ledgerDTO "github.com/eduardoths/finance-api/modules/ledger/dto"
)

type ExpenseService struct {
	repo   repositories.ExpenseRepository
	ledger gateways.LedgerGateway
}

func NewExpenseService(repo repositories.RepositoryContainer, gtw gateways.GatewayContainer) ExpenseService {
	return ExpenseService{}
}

func (es ExpenseService) Create(ctx context.Context, data dto.CreateExpense) (structs.Expense, error) {
	ts, err := es.ledger.NewTransaction(ctx, ledgerDTO.NewTransaction{
		Amount:      data.Amount,
		Date:        data.Date,
		Description: "Expense",
		AssetID:     data.Currency,
		CreditID:    data.FromAccount,
		DebitID:     domain.EXPENSE_ACCOUNT_ID,
	})
	if err != nil {
		return structs.Expense{}, err
	}

	expense, err := es.repo.Create(ctx, structs.Expense{
		Category:      data.Category,
		Amount:        data.Amount,
		Date:          data.Date,
		TransactionID: ts.ID,
		Description:   data.Description,
	})
	if err != nil {
		es.ledger.Rollback(ctx, ts.ID)
		return structs.Expense{}, err
	}

	if err := es.ledger.Commit(ctx, ts.ID); err != nil {
		// do later
	}

	return expense, nil
}
