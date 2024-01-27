package bookkeeper

import (
	"github.com/eduardoths/finance-api/modules/bookkeeper/controllers"
	"github.com/eduardoths/finance-api/modules/bookkeeper/gateways"
	"github.com/eduardoths/finance-api/modules/bookkeeper/repositories"
	"github.com/eduardoths/finance-api/modules/bookkeeper/services"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Route(fiber.Router)
}

type Bookkeeper struct {
	Controllers  []Controller
	services     services.ServiceContainer
	repositories repositories.RepositoryContainer
	gateways     gateways.GatewayContainer
}

func InitBookkeeper() Bookkeeper {
	bk := Bookkeeper{
		repositories: repositories.InitRepository(),
		gateways:     gateways.InitGateway(),
	}
	bk.services = services.InitService(bk.repositories, bk.gateways)
	bk.Controllers = []Controller{
		controllers.NewExpenseController(bk.services),
		controllers.NewTradeController(bk.services),
		controllers.NewIncomeController(bk.services),
	}
	return bk
}
