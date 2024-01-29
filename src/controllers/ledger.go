package controllers

import (
	"github.com/eduardoths/finance-api/modules/ledger"
	"github.com/gofiber/fiber/v2"
)

type LedgerController struct {
	ledger ledger.Ledger
}

func NewLedgerController() LedgerController {
	return LedgerController{
		ledger: ledger.NewLedger(),
	}
}

func (lc LedgerController) Route(r fiber.Router) {
	rg := r.Group("/ledger")
	for _, c := range lc.ledger.Controllers {
		c.Route(rg)
	}
}
