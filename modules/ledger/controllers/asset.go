package controllers

import (
	"net/http"

	"github.com/eduardoths/finance-api/modules/ledger/services"
	"github.com/eduardoths/finance-api/modules/ledger/structs"
	ihttp "github.com/eduardoths/finance-api/pkg/http"
	"github.com/gofiber/fiber/v2"
)

type AssetController struct {
	service services.AssetService
}

func NewAssetController(as services.AssetService) AssetController {
	return AssetController{
		service: as,
	}
}

func (ac AssetController) Create(c *fiber.Ctx) error {
	var body structs.Asset
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}
	data, err := ac.service.Create(c.UserContext(), body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(ihttp.Response[structs.Asset]{
		Data: data,
	})
}

func (ac AssetController) GetAll(c *fiber.Ctx) error {
	data, err := ac.service.GetAll(c.UserContext())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusCreated).JSON(ihttp.Response[[]structs.Asset]{
		Data: data,
	})
}

func (ac AssetController) Route(r fiber.Router) {
	rg := r.Group("/asset")

	rg.Post("/", ac.Create)
	rg.Get("/", ac.GetAll)
}
