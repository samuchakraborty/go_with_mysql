package controller

import (
	"fmt"
	"math"
	"mysql_with_go/database"
	"mysql_with_go/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllCustomer(c *fiber.Ctx) error {
	var customer []model.Customers
	rows := database.Database().Select("id, name, email, mobile").Find(&customer)

	fmt.Printf("%v", rows.RowsAffected)

	fmt.Printf("%v", rows.Error)

	listOfCustomer := []model.Customers{}

	fmt.Printf("%v\n", listOfCustomer)
	data := model.Data{
		StatusCode: 200,
		Message:    "Successfully get All customer",
		Count:      len(customer),
		Data:       customer,
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func InsertCustomer(c *fiber.Ctx) error {

	customer2 := new(model.Customers)

	if err := c.BodyParser(customer2); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	tx := database.Database().Begin()
	result := tx.FirstOrCreate(&customer2, model.Customers{Email: customer2.Email, Mobile: customer2.Mobile})

	if result.RowsAffected != 1 {
		tx.Rollback()
		return c.Status(422).JSON(fiber.Map{

			"msg":   "Duplicate data entry",
			"error": result.Error,
		})

	}
	tx.Commit()
	insertData := model.InsertData{
		StatusCode: fiber.StatusCreated,
		Message:    "Data Successfully Created",
		Data:       *customer2,
	}

	return c.Status(fiber.StatusCreated).JSON(insertData)

}

func GetAllCustomerByPagination(c *fiber.Ctx) error {
	var customers []model.Customers
	var total int64
	limit, _ := strconv.ParseInt(c.Query("page_size", "10"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page", "1"), 10, 64)
	if page == 0 {
		page = 1
	}

	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}

	offset := (page - 1) * limit

	// fmt.Printf("offset %v count %v\n", offset, total)

	database.Database().Find(&customers).Count(&total)

	database.Database().Select("id, name, email, mobile").Offset(int(offset)).Limit(int(limit)).Find(&customers)

	data := &model.CustomerPaginatedData{
		Current_page: page,
		Total_Record: total,
		Total_Page:   math.Ceil(float64(total) / float64(limit)),
		Customers:   customers,
	}

	return c.Status(200).JSON(data)

}
