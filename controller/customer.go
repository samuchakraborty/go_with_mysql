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

	rows, err := database.Database().Query("Select * from customers")

	if err != nil {
		log.Fatal(err)
	}
	var customer model.Customers
	listOfCustomer := []model.Customers{}

	for rows.Next() {

		err := rows.Scan(&customer.Id, &customer.Name, &customer.Mobile)

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

	insForm, err := database.Database().Prepare("INSERT INTO customers (name, mobile) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	res, err := insForm.Exec(customer2.Name, customer2.Mobile)

	defer database.Database().Close()
	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Last is %v", lastId)
	data := getCustomerByID(lastId)

	return c.JSON(data)

}

func GetCustomerByID(c *fiber.Ctx) error {

	result := c.Params("id")
	i, _ := strconv.ParseInt(result, 10, 64)
	data := getCustomerByID(i)

	return c.JSON(data)
}

func getCustomerByID(id int64) *model.InsertData {

	selectSql := "Select * from customers where id=" + strconv.FormatInt(id, 10)
	selectData, err := database.Database().Query(selectSql)

	if err != nil {
		log.Fatal(err)

	}
	var customer model.Customers

	for selectData.Next() {
		err := selectData.Scan(&customer.Id, &customer.Name, &customer.Mobile)

		if err != nil {
			log.Fatal(err)
		}
	}

	data := &model.InsertData{
		StatusCode: 200,
		Message:    "Successfully insert data on " + strconv.FormatInt(id, 10),
		Data:       customer,
	}

	return data

}
