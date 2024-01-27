package controllers

import (
	"net/http"

	"github.com/eduardoths/finance-api/modules/bookkeeper/dto"
	"github.com/eduardoths/finance-api/modules/bookkeeper/services"
	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	ihttp "github.com/eduardoths/finance-api/pkg/http"
	"github.com/gofiber/fiber/v2"
)

type ExpenseController struct {
	service services.ExpenseService
}

func NewExpenseController(servs services.ServiceContainer) ExpenseController {
	return ExpenseController{
		service: servs.Expense,
	}
}

func (ec ExpenseController) Create(c *fiber.Ctx) error {
	var body dto.CreateExpense
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	data, err := ec.service.Create(c.UserContext(), body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(ihttp.Response[structs.Expense]{
		Data: data,
	})
}

func (ec ExpenseController) Route(r fiber.Router) {
	rg := r.Group("/expense")

	rg.Post("/", ec.Create)
}
