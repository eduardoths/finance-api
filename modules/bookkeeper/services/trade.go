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

type TradeService struct {
	repo   repositories.TradeRepository
	ledger gateways.LedgerGateway
}

func NewTradeService(repos repositories.RepositoryContainer, gtw gateways.GatewayContainer) TradeService {
	return TradeService{
		repo:   repos.Trade,
		ledger: gtw.Ledger,
	}
}

func (ts TradeService) Create(ctx context.Context, data dto.CreateTrade) (structs.Trade, error) {
	sold, err := ts.ledger.NewTransaction(ctx, ledgerDTO.NewTransaction{
		Amount:      data.SoldAmount,
		Date:        data.Date,
		Description: "Trading assets",
		AssetID:     data.SoldAsset,
		CreditID:    data.FromAccount,
		DebitID:     domain.TRADE_ACCOUNT_ID,
	})
	if err != nil {
		return structs.Trade{}, err
	}

	fee, err := ts.ledger.NewTransaction(ctx, ledgerDTO.NewTransaction{
		Amount:      data.Fee,
		Date:        data.Date,
		Description: "Fee on trading assets",
		AssetID:     data.SoldAsset,
		CreditID:    data.FromAccount,
		DebitID:     domain.EXPENSE_ACCOUNT_ID,
	})
	if err != nil {
		return structs.Trade{}, err
	}

	purchased, err := ts.ledger.NewTransaction(ctx, ledgerDTO.NewTransaction{
		Amount:      data.PurchasedAmount,
		Date:        data.Date,
		Description: "Trading assets",
		AssetID:     data.PurchasedAsset,
		CreditID:    domain.TRADE_ACCOUNT_ID,
		DebitID:     data.ToAccount,
	})
	if err != nil {
		return structs.Trade{}, err
	}

	return ts.repo.Create(ctx, structs.Trade{
		TradeTransactions: []structs.TradeTransaction{
			{TransactionID: sold.ID},
			{TransactionID: fee.ID},
			{TransactionID: purchased.ID},
		},
		Fee:             data.Fee,
		Rate:            data.Rate,
		Date:            data.Date,
		SoldAmount:      data.SoldAmount,
		SoldAsset:       data.SoldAsset,
		PurchasedAmount: data.PurchasedAmount,
		PurchasedAsset:  data.PurchasedAsset,
	})
}

func (ts TradeService) GetAll(ctx context.Context) ([]dto.ViewTrade, error) {
	trades, err := ts.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]dto.ViewTrade, 0)

	for _, t := range trades {
		fromAsset, err := ts.ledger.GetAsset(ctx, t.SoldAsset)
		if err != nil {
			return nil, err
		}

		toAsset, err := ts.ledger.GetAsset(ctx, t.PurchasedAsset)
		if err != nil {
			return nil, err
		}

		data = append(data, dto.NewViewTrade(t, fromAsset, toAsset))
	}
	return data, nil
}
