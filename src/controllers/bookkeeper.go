package controllers

import (
	"github.com/eduardoths/finance-api/modules/bookkeeper"
	"github.com/gofiber/fiber/v2"
)

type BookkeeperController struct {
	bookkeeper bookkeeper.Bookkeeper
}

func NewBookkeeperController() BookkeeperController {
	return BookkeeperController{
		bookkeeper: bookkeeper.InitBookkeeper(),
	}
}

func (bc BookkeeperController) Route(r fiber.Router) {
	rg := r.Group("/bookkeeper")
	for _, c := range bc.bookkeeper.Controllers {
		c.Route(rg)
	}
}
