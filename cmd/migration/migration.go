package main

import (
	"context"

	"github.com/eduardoths/finance-api/modules/ledger/domain"
	"github.com/eduardoths/finance-api/pkg/db"
)

func main() {
	db.StartDB()
	defer db.CloseConn()

	conn := db.GetFromContext(context.Background())
	conn.AutoMigrate(domain.Tables...)

	conn.Create(&accounts)
	conn.Create(&assets)
}
