package controllers

import (
	"net/http"

	"github.com/eduardoths/finance-api/modules/bookkeeper/dto"
	"github.com/eduardoths/finance-api/modules/bookkeeper/services"
	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	ihttp "github.com/eduardoths/finance-api/pkg/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type TradeController struct {
	service services.TradeService
}

func NewTradeController(servs services.ServiceContainer) TradeController {
	return TradeController{
		service: servs.Trade,
	}
}

func (tc TradeController) Create(c *fiber.Ctx) error {
	var body dto.CreateTrade
	if err := c.BodyParser(&body); err != nil {
		log.Error(err)
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	data, err := tc.service.Create(c.UserContext(), body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(ihttp.Response[structs.Trade]{
		Data: data,
	})
}

func (tc TradeController) Route(r fiber.Router) {
	rg := r.Group("/trade")

	rg.Post("/", tc.Create)
}
