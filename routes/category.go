package routes

import (
	"errors"
	"qr-menu-api/config"
	"qr-menu-api/models"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func CreateResponseCategory(category models.Category) Category {
	return Category{ID: category.ID, Name: category.Name}
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	config.Database.Db.Create(&category)

	responseCategory := CreateResponseCategory(category)

	return c.Status(200).JSON(responseCategory)
}

func GetCategories(c *fiber.Ctx) error {
	categories := []models.Category{}
	config.Database.Db.Find(&categories)

	responseCategories := []Category{}
	for _, category := range categories {
		responseCategory := CreateResponseCategory(category)

		responseCategories = append(responseCategories, responseCategory)
	}

	return c.Status(200).JSON(responseCategories)
}

func findCategory(id int, category *models.Category) error {
	config.Database.Db.Find(&category, "id=?", id)
	if category.ID == 0 {
		return errors.New("Category does not exist")
	}
	return nil
}

func GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findCategory(id, &category); err != nil {
		c.Status(400).JSON(err.Error())
	}

	responseCategory := CreateResponseCategory(category)
	return c.Status(200).JSON(responseCategory)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findCategory(id, &category); err != nil {
		c.Status(400).JSON(err.Error())
	}

	type UpdateCategory struct {
		Name string `json:"name"`
	}

	var updatedata UpdateCategory

	category.Name = updatedata.Name

	config.Database.Db.Save(&category)

	responseCategory := CreateResponseCategory(category)

	return c.Status(200).JSON(responseCategory)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findCategory(id, &category); err != nil {
		c.Status(400).JSON(err.Error())
	}

	if err := config.Database.Db.Delete(&category).Error; err != nil {
		c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Category deleted successfully")
}
