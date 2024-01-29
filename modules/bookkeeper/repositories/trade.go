package repositories

import (
	"context"

	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	"github.com/eduardoths/finance-api/pkg/db"
)

type TradeRepository struct{}

func NewTradeRepository() TradeRepository {
	return TradeRepository{}
}

func (tr TradeRepository) Create(ctx context.Context, trade structs.Trade) (structs.Trade, error) {
	if err := db.GetFromContext(ctx).
		Create(&trade).
		Error; err != nil {
		return structs.Trade{}, err
	}
	return trade, nil
}

func (tr TradeRepository) GetAll(ctx context.Context) ([]structs.Trade, error) {
	var data []structs.Trade
	if err := db.GetFromContext(ctx).
		Find(&data).
		Error; err != nil {
		return nil, err
	}
	return data, nil
}
