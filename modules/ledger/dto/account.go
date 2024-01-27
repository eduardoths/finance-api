package dto

import "github.com/eduardoths/finance-api/modules/ledger/structs"

type NewAccount struct {
	Name string
	Type structs.AccountType
}
