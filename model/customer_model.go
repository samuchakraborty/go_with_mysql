package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Customers struct {
	*gorm.Model
	Id     uint   `gorm:"primaryKey, autoIncrement" json:"id" query:"id" form:"id"`
	Name   string `validate:"required,min=3,max=32" json:"name" query:"name" form:"name"`
	Mobile string `gorm:"unique" validate:"required,max=11" json:"mobile" query:"mobile" form:"mobile"`
	Email  string `gorm:"unique" validate:"required,email" json:"email" query:"email" form:"email"`
}
type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func ValidateAddCustomer(c *fiber.Ctx) error {
	var errors []*IError
	body := new(Customers)
	c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Error()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return c.Next()
}

type CustomerPaginatedData struct {
	Total_Record int64 `json:"total_record"`
	Current_page int64 `json:"current_Page"`
	// Count        int `json:"count"`
	Total_Page float64 `json:"total_page"`

	Customers []Customers `json:"customerData"`
}
type User struct {
	*gorm.Model
	UserName  string  `json:"userName"`
	CompanyID int     `json:"companyId"`
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	ID             int    `gorm:"primaryKey, autoIncrement" json:"id" query:"id" form:"id"`
	CompanyName    string `gorm:"unique" json:"companyName"`
	CompanyAddress string `gorm:"unique" json:"companyAddress"`
}
