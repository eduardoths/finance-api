package domain

import "github.com/eduardoths/finance-api/modules/ledger/structs"

var Tables []interface{} = []interface{}{
	&structs.Account{},
	&structs.Asset{},
	&structs.Transaction{},
}
