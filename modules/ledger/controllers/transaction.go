package controllers

import (
	"net/http"

	"github.com/eduardoths/finance-api/modules/ledger/dto"
	"github.com/eduardoths/finance-api/modules/ledger/services"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
	ihttp "github.com/eduardoths/finance-api/pkg/http"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(ts services.TransactionService) TransactionController {
	return TransactionController{
		service: ts,
	}
}

func (tc TransactionController) Create(c *fiber.Ctx) error {
	var body dto.NewTransaction
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	data, err := tc.service.New(c.UserContext(), body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(ihttp.Response[structs.Transaction]{
		Data: data,
	})
}

func (tc TransactionController) GetAll(c *fiber.Ctx) error {
	data, err := tc.service.GetAll(c.UserContext())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusCreated).JSON(ihttp.Response[[]structs.Transaction]{
		Data: data,
	})
}

func (tc TransactionController) Route(r fiber.Router) {
	rg := r.Group("/transaction")

	rg.Post("/", tc.Create)
	rg.Get("/", tc.GetAll)
}
