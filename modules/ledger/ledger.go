package ledger

import (
	"github.com/eduardoths/finance-api/modules/ledger/controllers"
	"github.com/eduardoths/finance-api/modules/ledger/repositories"
	"github.com/eduardoths/finance-api/modules/ledger/services"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Route(fiber.Router)
}

type Ledger struct {
	Accounts     services.AccountService
	Transactions services.TransactionService
	Assets       services.AssetService
	Controllers  []Controller
}

func NewLedger() Ledger {
	accountRepo := repositories.AccountRepository{}
	transactionRepo := repositories.TransactionRepository{}
	transactionService := services.NewTransactionService(transactionRepo)
	assetService := services.NewAssetService(repositories.AssetRepository{})
	accountService := services.NewAccountService(accountRepo, transactionService, assetService)
	return Ledger{
		Transactions: transactionService,
		Accounts:     accountService,
		Assets:       assetService,
		Controllers: []Controller{
			controllers.NewAccountController(accountService),
			controllers.NewAssetController(assetService),
			controllers.NewTransactionController(transactionService),
		},
	}
}
