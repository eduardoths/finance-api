package structs

import "time"

type Expense struct {
	Model
	Category      ExpenseCategory
	Amount        uint
	Date          time.Time
	TransactionID ExternalID
	Description   string
}
