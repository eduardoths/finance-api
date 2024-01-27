package repositories

type RepositoryContainer struct {
	Expense ExpenseRepository
	Trade   TradeRepository
	Income  IncomeRepository
}

func InitRepository() RepositoryContainer {
	return RepositoryContainer{
		Expense: NewExpenseRepository(),
		Trade:   NewTradeRepository(),
	}
}
