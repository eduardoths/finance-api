package controllers

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Route(r fiber.Router)
}
type Controllers []Controller

func StartControllers(router fiber.Router) Controllers {
	ctrl := Controllers{
		NewLedgerController(),
		NewBookkeeperController(),
	}
	for _, c := range ctrl {
		c.Route(router)
	}
	return ctrl
}
