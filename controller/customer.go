package controller

import (
	"fmt"
	"log"
	"mysql_with_go/database"
	"mysql_with_go/model"
	"strconv"

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
	fmt.Println(string(c.Request().URI().QueryString()))

	customer2 := new(model.Customers)
	if err := c.QueryParser(customer2); err != nil {
		return err
	}

	fmt.Printf("%v %v %v\n", customer2, customer2.Fullname, customer2.Mobile)
	sql := "INSERT INTO customer VALUES ( 0," + customer2.Fullname + "," + customer2.Mobile + ")"
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

	selectSql := "Select * from customer where id=" + strconv.FormatInt(lastId, 10)
	selectData, err := database.Database().Query(selectSql)

	if err != nil {
		log.Fatal(err)

	}
	var customer model.Customers

	for selectData.Next() {
		err := selectData.Scan(&customer.Id, &customer.Fullname, &customer.Mobile)

		if err != nil {
			log.Fatal(err)
		}
	}

	data := &model.InsertData{
		StatusCode: 200,
		Message:    "Successfully insert data on " + strconv.FormatInt(lastId, 10),
		Data:       customer,
	}

	return c.JSON(data)

}
