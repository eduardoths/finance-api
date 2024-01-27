package dto

import (
	"time"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
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
