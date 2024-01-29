package ledger

import (
	"github.com/eduardoths/finance-api/modules/ledger/repositories"
	"github.com/eduardoths/finance-api/modules/ledger/services"
)

type Ledger struct {
	Accounts     services.AccountService
	Transactions services.TransactionService
	Assets       services.AssetService
}

func NewLedger() Ledger {
	accountRepo := repositories.AccountRepository{}
	transactionRepo := repositories.TransactionRepository{}
	transactionService := services.NewTransactionService(transactionRepo)
	assetService := services.NewAssetService(repositories.AssetRepository{})
	return Ledger{
		Transactions: transactionService,
		Accounts:     services.NewAccountService(accountRepo, transactionService, assetService),
		Assets:       assetService,
	}
}
