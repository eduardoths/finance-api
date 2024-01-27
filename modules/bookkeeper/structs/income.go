package structs

import "time"

type Income struct {
	Model
	Category      IncomeCategory
	Amount        uint
	Date          time.Time
	TransactionID ExternalID
	Description   string
}
