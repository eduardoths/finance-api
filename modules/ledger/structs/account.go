package structs

type Account struct {
	Model
	Name string
	Type AccountType
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
	DEBIT_NORMAL  AccountType = "DEBIT"
	CREDIT_NORMAL AccountType = "CREDIT"
)
