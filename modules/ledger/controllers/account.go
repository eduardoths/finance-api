package controllers

import (
	"net/http"

	"github.com/eduardoths/finance-api/modules/ledger/dto"
	"github.com/eduardoths/finance-api/modules/ledger/services"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
	ihttp "github.com/eduardoths/finance-api/pkg/http"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	service services.AccountService
}

func NewAccountController(as services.AccountService) AccountController {
	return AccountController{
		service: as,
	}
}

func (ac AccountController) Create(c *fiber.Ctx) error {
	var body dto.NewAccount
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	data, err := ac.service.Create(c.UserContext(), body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(ihttp.Response[structs.Account]{
		Data: data,
	})
}

func (ac AccountController) GetAll(c *fiber.Ctx) error {
	data, err := ac.service.GetAll(c.UserContext())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusCreated).JSON(ihttp.Response[[]structs.Account]{
		Data: data,
	})
}

func (ac AccountController) Route(r fiber.Router) {
	rg := r.Group("/account")

	rg.Post("/", ac.Create)
	rg.Get("/", ac.GetAll)
}
