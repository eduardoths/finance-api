package dto

import (
	"time"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	ledgerStructs "github.com/eduardoths/finance-api/modules/ledger/structs"
)

type CreateTrade struct {
	FromAccount     structs.ExternalID `json:"from_account"`
	ToAccount       structs.ExternalID `json:"to_account"`
	Date            time.Time          `json:"date"`
	Fee             uint               `json:"fee"`
	Rate            uint               `json:"rate"`
	SoldAmount      uint               `json:"sold_amount"`
	SoldAsset       structs.ExternalID `json:"sold_asset"`
	PurchasedAmount uint               `json:"purchased_amount"`
	PurchasedAsset  structs.ExternalID `json:"purchased_asset"`
}

type ViewTrade struct {
	ID              uint                `json:"id"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
	Date            time.Time           `json:"date"`
	Fee             uint                `json:"fee"`
	Rate            uint                `json:"rate"`
	SoldAmount      uint                `json:"sold_amount"`
	SoldAsset       ledgerStructs.Asset `json:"sold_asset"`
	PurchasedAmount uint                `json:"purchased_amount"`
	PurchasedAsset  ledgerStructs.Asset `json:"purchased_asset"`
}

func NewViewTrade(t structs.Trade, soldAsset, purchasedAsset ledgerStructs.Asset) ViewTrade {
	return ViewTrade{
		ID:              t.ID,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		Date:            t.Date,
		Fee:             t.Fee,
		Rate:            t.Rate,
		SoldAsset:       soldAsset,
		SoldAmount:      t.SoldAmount,
		PurchasedAsset:  purchasedAsset,
		PurchasedAmount: t.PurchasedAmount,
	}
}
