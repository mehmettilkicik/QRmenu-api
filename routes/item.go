package routes

import (
	"errors"
	"qr-menu-api/config"
	"qr-menu-api/models"

	"github.com/gofiber/fiber/v2"
)

/*
{
	id:1,
	name:"drinks"
}
*/

type Item struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

func CreateResponseItem(item models.Item, category Category) Item {
	return Item{ID: item.ID, Name: item.Name, Category: category}
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var category models.Category
	if err := findCategory(item.CategoryRefer, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	config.Database.Db.Create(&item)

	responseCategory := CreateResponseCategory(category)
	responseItem := CreateResponseItem(item, responseCategory)

	return c.Status(200).JSON(responseItem)
}

func GetItems(c *fiber.Ctx) error {
	items := []models.Item{}
	config.Database.Db.Find(&items)

	responseItems := []Item{}

	for _, item := range items {
		var category models.Category
		config.Database.Db.Find(&category, "id=?", item.CategoryRefer)

		responseItem := CreateResponseItem(item, CreateResponseCategory(category))
		responseItems = append(responseItems, responseItem)
	}

	return c.Status(200).JSON(responseItems)
}

func findItemsByCategory(cRefer int, item *models.Item) error {
	config.Database.Db.Find(&item, "category_id=?", cRefer)
	if item.CategoryRefer == 0 {
		return errors.New("no items in this category")
	}
	return nil
}

func findItem(id int, item *models.Item) error {
	config.Database.Db.Find(&item, "id=?", id)
	if item.CategoryRefer == 0 {
		return errors.New("item does not exist")
	}
	return nil
}

func GetItem(c *fiber.Ctx) error {
	return c.Status(200).SendString("return")
}
