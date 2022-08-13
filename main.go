package main

import (
	"mysql_with_go/controller"
	"mysql_with_go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get(routes.GetAllCustomer, controller.GetAllCustomer)

	app.Get(routes.InsertCustomer, controller.InsertCustomer)

	app.Listen(":3000")

}
