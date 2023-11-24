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

//Get item by category

func findItemsByCategory(cRefer int, items *[]models.Item) error {
	config.Database.Db.Find(&items, "category_refer=?", cRefer)
	if len(*items) == 0 {
		return errors.New("no items in this category")
	}
	return nil
}

func GetItemsByCategory(c *fiber.Ctx) error {
	cRefer, err := c.ParamsInt("category_refer")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var items []models.Item
	if err := findItemsByCategory(cRefer, &items); err != nil {
		if err.Error() == "no items in this category" {
			return c.Status(404).JSON("No items found in this category")
		}
		return c.Status(500).JSON("Internal server errror")
	}

	responseItems := []Item{}
	for _, item := range items {
		var category models.Category
		config.Database.Db.Find(&category, "id=?", item.CategoryRefer)
		responseCategory := CreateResponseCategory(category)
		responseItem := CreateResponseItem(item, responseCategory)
		responseItems = append(responseItems, responseItem)
	}
	return c.Status(200).JSON(responseItems)
}

//Get Specific Item

func findItem(cRefer int, id int, item *models.Item) error {
	config.Database.Db.Where("category_refer=?", cRefer).Find(&item, "id=?", id)
	if item.ID == 0 {
		return errors.New("item does not exist")
	}
	return nil
}

func GetItem(c *fiber.Ctx) error {
	cRefer, err := c.ParamsInt("category_refer")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :category_refer is an integer")
	}
	id, err2 := c.ParamsInt("id")
	if err2 != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	var item models.Item
	if err := findItem(cRefer, id, &item); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var category models.Category
	if err := findCategory(cRefer, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseCategory := CreateResponseCategory(category)

	responseItem := CreateResponseItem(item, responseCategory)

	return c.Status(200).JSON(responseItem)

}

//Update Item

//Delete Item
