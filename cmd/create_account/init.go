package main

import (
	"context"
	"fmt"

	"github.com/eduardoths/finance-api/modules/ledger"
	"github.com/eduardoths/finance-api/modules/ledger/dto"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
	"github.com/eduardoths/finance-api/pkg/db"
)

func main() {

	db.StartDB()
	defer db.CloseConn()
	ctx := context.Background()

	l := ledger.NewLedger()

	// l.Transactions.New(ctx, dto.NewTransaction{
	// 	Amount:      100_00,
	// 	Date:        time.Now(),
	// 	Description: "Received 100 reals",
	// 	AssetID:     1,
	// 	DebitID:     2,
	// 	CreditID:    1,
	// })

	// l.Transactions.New(ctx, dto.NewTransaction{
	// 	Amount:      100_00,
	// 	Date:        time.Now(),
	// 	Description: "Convert real to eur",
	// 	AssetID:     1,
	// 	DebitID:     1,
	// 	CreditID:    2,
	// })

	// l.Transactions.New(ctx, dto.NewTransaction{
	// 	Amount:      10_00,
	// 	Date:        time.Now(),
	// 	Description: "Convert real to eur",
	// 	AssetID:     2,
	// 	DebitID:     3,
	// 	CreditID:    1,
	// })

	b1, _ := l.Accounts.CalculateBalance(ctx, 1)
	b2, _ := l.Accounts.CalculateBalance(ctx, 2)
	b3, _ := l.Accounts.CalculateBalance(ctx, 3)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
}

func initData(ctx context.Context, l ledger.Ledger) {
	accounts := []dto.NewAccount{
		{
			Name: "world",
			Type: structs.DEBIT_NORMAL,
		},
		{
			Name: "Wise-BRL",
			Type: structs.DEBIT_NORMAL,
		},
		{
			Name: "Wise-EUR",
			Type: structs.DEBIT_NORMAL,
		},
		{
			Name: "Santander-PT",
			Type: structs.DEBIT_NORMAL,
		},
	}

	assets := []structs.Asset{
		{
			Type: structs.CURRENCY,
			Code: "BRL",
			Name: "Real",
		},
		{
			Type: structs.CURRENCY,
			Code: "EUR",
			Name: "Euro",
		},
		{
			Type: structs.CURRENCY,
			Code: "USD",
			Name: "Dollar",
		},
	}

	for _, a := range accounts {
		l.Accounts.Create(ctx, a)
	}

	for _, a := range assets {
		db.GetFromContext(ctx).Create(&a)
	}
}
