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

type IncomeService struct {
	repo   repositories.IncomeRepository
	ledger gateways.LedgerGateway
}

func NewIncomeService(repo repositories.RepositoryContainer, gtw gateways.GatewayContainer) IncomeService {
	return IncomeService{}
}

func (is IncomeService) Create(ctx context.Context, data dto.CreateIncome) (structs.Income, error) {
	ts, err := is.ledger.NewTransaction(ctx, ledgerDTO.NewTransaction{
		Amount:      data.Amount,
		Date:        data.Date,
		Description: "Income",
		AssetID:     data.Currency,
		CreditID:    domain.INCOME_ACCOUNT_ID,
		DebitID:     data.ToAccount,
	})
	if err != nil {
		return structs.Income{}, err
	}

	income, err := is.repo.Create(ctx, structs.Income{
		Category:      data.Category,
		Amount:        data.Amount,
		Date:          data.Date,
		TransactionID: ts.ID,
		Description:   data.Description,
	})
	if err != nil {
		is.ledger.Rollback(ctx, ts.ID)
		return structs.Income{}, err
	}

	if err := is.ledger.Commit(ctx, ts.ID); err != nil {
		// do later
	}

	return income, nil
}
