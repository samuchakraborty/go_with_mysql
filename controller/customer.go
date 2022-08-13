package controller

import (
	"fmt"
	"log"
	"mysql_with_go/database"
	"mysql_with_go/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomer(c *fiber.Ctx) error {

	rows, err := database.Database().Query("Select * from customer")

	if err != nil {
		log.Fatal(err)
	}
	var customer model.Customers
	listOfCustomer := []model.Customers{}

	for rows.Next() {

		err := rows.Scan(&customer.Id, &customer.Fullname, &customer.Mobile)

		if err != nil {
			log.Fatal(err)
		}
		listOfCustomer = append(listOfCustomer, customer)

	}
	fmt.Printf("%v\n", listOfCustomer)
	data := model.Data{
		StatusCode: 200,
		Message:    "Successfully get All customer",
		Data:       listOfCustomer,
	}

	return c.JSON(data)
}

func InsertCustomer(c *fiber.Ctx) error {

	sql := "INSERT INTO customer VALUES( 0,'saas', 01091944605)"
	fmt.Printf("%v\n", sql)

	res, err := database.Database().Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Last is %v", lastId)
	return c.SendString("insert Successfully")

}
