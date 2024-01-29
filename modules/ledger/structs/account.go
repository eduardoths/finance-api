package structs

type Account struct {
	Model
	Name    string           `json:"name"`
	Type    AccountType      `json:"type"`
	Balance []AccountBalance `json:"balance" gorm:"-"`
}

type AccountBalance struct {
	Asset   Asset
	Balance int
}

func (a Account) DebitMultiplier() int {
	if a.Type == DEBIT_NORMAL {
		return 1
	}
	return -1
}

func (a Account) CreditMultiplier() int {
	if a.Type == CREDIT_NORMAL {
		return 1
	}
	return -1
}

type AccountType string

const (
	DEBIT_NORMAL  AccountType = "DEBIT_NORMAL"
	CREDIT_NORMAL AccountType = "CREDIT_NORMAL"
)
