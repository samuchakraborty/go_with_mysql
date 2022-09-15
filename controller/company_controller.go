package controller

import (
	"mysql_with_go/database"
	"mysql_with_go/model"

	"github.com/gofiber/fiber/v2"
)

func AddCompany(c *fiber.Ctx) error {
	company := new(model.Company)

	err := c.BodyParser(company)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx := database.Database().Begin()

	result := tx.FirstOrCreate(&company,
		model.Company{CompanyName: company.CompanyName, CompanyAddress: company.CompanyAddress})

	if result.RowsAffected != 1 {

		tx.Rollback()
		return c.Status(422).JSON(fiber.Map{
			"msg":   "Duplicate data entry",
			"error": result.Error,
		})

	}
	tx.Commit()

	insertData := model.CompanyInsertData{
		StatusCode: fiber.StatusCreated,
		Message:    "Data Successfully Created",
		Data:       *company,
	}

	return c.Status(200).JSON(insertData)

}
