package routes

import (
	"errors"
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

func GetTables(c *fiber.Ctx) error {
	tables := []models.Table{}

	config.Database.Db.Find(&tables)
	responseTables := []Table{}

	for _, table := range tables {
		responseTable := CreateResponseTable(table)
		responseTables = append(responseTables, responseTable)
	}

	return c.Status(200).JSON(responseTables)
}

func findTable(id int, table *models.Table) error {
	config.Database.Db.Find(&table, "id=?", id)
	if table.ID == 0 {
		return errors.New("table does not exist")
	}
	return nil
}

func GetTable(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var table models.Table

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findTable(id, &table); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseTable := CreateResponseTable(table)

	return c.Status(200).JSON(responseTable)
}

func UpdateTable(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var table models.Table

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findTable(id, &table); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateTable struct {
		Number uint `json:"number"`
	}

	var updatedata UpdateTable

	if err := c.BodyParser(&updatedata); err != nil {
		c.Status(500).JSON(err.Error())
	}

	table.Number = updatedata.Number

	config.Database.Db.Save(&table)

	responseTable := CreateResponseTable(table)

	return c.Status(200).JSON(responseTable)
}

func DeleteTable(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var table models.Table

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findTable(id, &table); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := config.Database.Db.Delete(&table).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Table deleted successfully")
}
