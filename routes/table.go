package routes

import (
	"qr-menu-api/config"
	"qr-menu-api/models"

	"github.com/gofiber/fiber/v2"
)

type Table struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	Number uint `json:"number"`
}

func CreateResponseTable(table models.Table) Table {
	return Table{ID: table.ID, Number: table.Number}
}

func CreateTable(c *fiber.Ctx) error {
	var table models.Table

	if err := c.BodyParser(&table); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	config.Database.Db.Create(&table)
	responseTable := CreateResponseTable(table)
	return c.Status(200).JSON(responseTable)
}
