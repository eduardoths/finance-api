package controllers

import (
	"net/http"

	"github.com/eduardoths/finance-api/modules/bookkeeper/dto"
	"github.com/eduardoths/finance-api/modules/bookkeeper/services"
	"github.com/eduardoths/finance-api/modules/bookkeeper/structs"
	ihttp "github.com/eduardoths/finance-api/pkg/http"
	"github.com/gofiber/fiber/v2"
)

type IncomeController struct {
	service services.IncomeService
}

func NewIncomeController(servs services.ServiceContainer) IncomeController {
	return IncomeController{
		service: servs.Income,
	}
}

func (ec IncomeController) Create(c *fiber.Ctx) error {
	var body dto.CreateIncome
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	data, err := ec.service.Create(c.UserContext(), body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(ihttp.Response[structs.Income]{
		Data: data,
	})
}

func (ec IncomeController) Route(r fiber.Router) {
	rg := r.Group("/income")

	rg.Post("/", ec.Create)
}
