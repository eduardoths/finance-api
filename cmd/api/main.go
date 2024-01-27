package main

import (
	"github.com/eduardoths/finance-api/pkg/db"
	"github.com/eduardoths/finance-api/src/controllers"
	"github.com/eduardoths/finance-api/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.StartDB()
	defer db.CloseConn()

	app := fiber.New()
	router := app.Group("/api/v1/").
		Use(middlewares.WithContext())

	controllers.StartControllers(router)
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
