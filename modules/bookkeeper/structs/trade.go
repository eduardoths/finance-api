package structs

import "time"

type Trade struct {
	Model
	TradeTransactions []TradeTransaction
	Fee               uint
	Rate              uint
	Date              time.Time

	SoldAmount      uint
	SoldAsset       ExternalID
	PurchasedAmount uint
	PurchasedAsset  ExternalID
}

type TradeTransaction struct {
	Model
	TransactionID []ExternalID
	TradeID       ID
}
