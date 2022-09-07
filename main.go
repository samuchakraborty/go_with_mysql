package main

import (
	"mysql_with_go/controller"
	"mysql_with_go/database"
	"mysql_with_go/model"
	"mysql_with_go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.Database().AutoMigrate(&model.Customers{})
	database.Database().AutoMigrate(&model.Company{})
	database.Database().AutoMigrate(&model.User{})

	database.Database().AutoMigrate(&model.Company{})


	database.Database().Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&model.Customers{})
	app.Get(routes.GetAllCustomer, controller.GetAllCustomer)

	app.Post(routes.InsertCustomer, model.ValidateAddCustomer, controller.InsertCustomer)

	app.Get(routes.GetAllCustomerByPaginated, controller.GetAllCustomerByPagination)
	// app.Get(routes.GetCustomerById, controller.GetCustomerByID)
	// app.Delete(routes.DeleteCustomerById, controller.DeleteCustomerByID)

	// app.Patch(routes.UpDateCustomerById, controller.UpdateCustomerByID)

	app.Listen(":3000")



	

}
