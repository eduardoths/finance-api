package services

import (
	"github.com/eduardoths/finance-api/modules/bookkeeper/gateways"
	"github.com/eduardoths/finance-api/modules/bookkeeper/repositories"
)

type ServiceContainer struct {
	Expense ExpenseService
	Trade   TradeService
	Income  IncomeService
}

func InitService(repo repositories.RepositoryContainer, gtw gateways.GatewayContainer) ServiceContainer {
	return ServiceContainer{
		Expense: NewExpenseService(repo, gtw),
		Trade:   NewTradeService(repo, gtw),
		Income:  NewIncomeService(repo, gtw),
	}
}
