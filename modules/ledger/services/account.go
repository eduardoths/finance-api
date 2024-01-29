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
	assetService       AssetService
}

func NewAccountService(repo repositories.AccountRepository, ts TransactionService, as AssetService) AccountService {
	return AccountService{
		repo:               repo,
		transactionService: ts,
		assetService:       as,
	}
}

func (as AccountService) GetAll(ctx context.Context) ([]structs.Account, error) {
	return as.repo.GetAll(ctx)
}

func (as AccountService) Create(ctx context.Context, data dto.NewAccount) (structs.Account, error) {
	return as.repo.Create(ctx, structs.Account{Name: data.Name, Type: data.Type})
}

func (as AccountService) Get(ctx context.Context, id structs.ID) (structs.Account, error) {
	account, err := as.repo.Get(ctx, id)
	if err != nil {
		return structs.Account{}, err
	}
	balance, err := as.calculateBalance(ctx, id)
	if err != nil {
		return structs.Account{}, err
	}
	account.Balance = balance
	return account, nil
}

func (as AccountService) calculateBalance(ctx context.Context, accountID structs.ID) ([]structs.AccountBalance, error) {
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

	balanceMap := make(map[structs.ID]int)
	for _, t := range ts {
		if t.DebitID == accountID {
			balanceMap[t.AssetID] += int(t.Amount) * account.DebitMultiplier()
		} else {
			balanceMap[t.AssetID] += int(t.Amount) * account.CreditMultiplier()
		}
	}

	balance := make([]structs.AccountBalance, 0)
	for k, v := range balanceMap {
		asset, err := as.assetService.Get(ctx, k)
		if err != nil {
			continue
		}
		balance = append(balance, structs.AccountBalance{
			Balance: v,
			Asset:   asset,
		})
	}

	return balance, nil
}
