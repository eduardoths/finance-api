package services

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger/dto"
	"github.com/eduardoths/finance-api/modules/ledger/repositories"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
)

type AccountService struct {
	repo               repositories.AccountRepository
	transactionService TransactionService
}

func NewAccountService(repo repositories.AccountRepository, ts TransactionService) AccountService {
	return AccountService{
		repo:               repo,
		transactionService: ts,
	}
}

func (as AccountService) Create(ctx context.Context, data dto.NewAccount) (structs.Account, error) {
	return as.repo.Create(ctx, structs.Account{Name: data.Name, Type: data.Type})
}

func (as AccountService) Get(ctx context.Context, id structs.ID) (structs.Account, error) {
	return as.repo.Get(ctx, id)
}

func (as AccountService) CalculateBalance(ctx context.Context, accountID structs.ID) (map[structs.ID]int, error) {
	account, err := as.repo.Get(ctx, accountID)
	if err != nil {
		return nil, err
	}

	ts, err := as.transactionService.GetMany(ctx, dto.TransactionFilter{
		AccountId: &accountID,
	})
	if err != nil {
		return nil, err
	}

	balance := make(map[structs.ID]int)
	for _, t := range ts {
		if t.DebitID == accountID {
			balance[t.AssetID] += int(t.Amount) * account.DebitMultiplier()
		} else {
			balance[t.AssetID] += int(t.Amount) * account.CreditMultiplier()
		}
	}

	return balance, nil
}
