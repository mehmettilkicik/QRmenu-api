package routes

import (
	"errors"
	"qr-menu-api/config"
	"qr-menu-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Table        Table         `gorm:"foreignKey:TableRefer"`
	OrderDetails []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID"`
	IsPaid       bool          `json:"is_paid"`
}

type OrderDetail struct {
	OrderID  uint `json:"order_id" gorm:"index"`
	Item     Item
	Quantity uint `json:"quantity"`
}

func CreateResponseOrder(order models.Order, table Table, ordedetails []OrderDetail) Order {
	return Order{ID: order.ID, CreatedAt: order.CreatedAt, Table: table, OrderDetails: ordedetails}
}

func CreateResponseOrderDetail(ordedetail models.OrderDetail, item Item) OrderDetail {
	return OrderDetail{OrderID: ordedetail.OrderID, Item: item, Quantity: ordedetail.Quantity}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var table models.Table

	if err := findTable(order.TableRefer, &table); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseDetails := []OrderDetail{}
	config.Database.Db.Create(&order)
	for _, v := range order.OrderDetails {
		v.OrderID = order.ID
		var item models.Item
		if err := findItemByID(v.ItemRefer, &item); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		var category models.Category
		if err := findCategory(item.CategoryRefer, &category); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responseCategory := CreateResponseCategory(category)
		responseItem := CreateResponseItem(item, responseCategory)
		responseDetails = append(responseDetails, CreateResponseOrderDetail(v, responseItem))
	}
	responseTable := CreateResponseTable(table)
	responseOrder := CreateResponseOrder(order, responseTable, responseDetails)

	return c.Status(200).JSON(responseOrder)
}

// Find All orders
func findOrders(orders *[]models.Order) error {
	config.Database.Db.Find(&orders)
	if len(*orders) == 0 {
		return errors.New("there is no order")
	}
	return nil
}

// Get All orders
func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	if err := findOrders(&orders); err != nil {
		return c.Status(400).JSON("error")
	}

	responseOrders := []Order{}
	responseDetails := []OrderDetail{}
	for _, order := range orders {
		var table models.Table
		config.Database.Db.Find(&table, "id=?", order.TableRefer)
		var orderDetails []models.OrderDetail
		config.Database.Db.Find(&orderDetails, "order_id=?", order.ID)
		for _, v := range orderDetails {
			var item models.Item
			if err := findItemByID(v.ItemRefer, &item); err != nil {
				return c.Status(400).JSON(err.Error())
			}
			var category models.Category
			if err := findCategory(item.CategoryRefer, &category); err != nil {
				return c.Status(400).JSON(err.Error())
			}
			responseCategory := CreateResponseCategory(category)
			responseItem := CreateResponseItem(item, responseCategory)
			responseDetails = append(responseDetails, CreateResponseOrderDetail(v, responseItem))

		}
		responseTable := CreateResponseTable(table)
		responseOrder := CreateResponseOrder(order, responseTable, responseDetails)
		responseOrders = append(responseOrders, responseOrder)
		responseDetails = []OrderDetail{}
	}
	return c.Status(200).JSON(responseOrders)
}

//

func findOrdersByTable(tableRefer int, orders *[]models.Order) error {
	config.Database.Db.Find(&orders, "table_refer=?", tableRefer)
	if len(*orders) == 0 {
		return errors.New("no orders in this table")
	}
	return nil
}

func findOrderByTable(tableRefer int, order *models.Order, isPaid int) error {
	config.Database.Db.Where("table_refer=?", tableRefer).Find(&order, "is_paid=?", isPaid)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetAllOrders(c *fiber.Ctx) error {
	tableRefer, err := c.ParamsInt("table_refer")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :table_refer is an integer")
	}

	var orders []models.Order
	if err := findOrdersByTable(tableRefer, &orders); err != nil {
		if err.Error() == "no orders in this table" {
			return c.Status(404).JSON("No items found in this table")
		}
		return c.Status(500).JSON("Internal server errror")
	}

	responseOrders := []Order{}
	responseDetails := []OrderDetail{}
	for _, order := range orders {
		var table models.Table
		config.Database.Db.Find(&table, "id=?", order.TableRefer)
		var orderDetails []models.OrderDetail
		config.Database.Db.Find(&orderDetails, "order_id=?", order.ID)
		for _, v := range orderDetails {
			var item models.Item
			if err := findItemByID(v.ItemRefer, &item); err != nil {
				return c.Status(400).JSON(err.Error())
			}
			var category models.Category
			if err := findCategory(item.CategoryRefer, &category); err != nil {
				return c.Status(400).JSON(err.Error())
			}
			responseCategory := CreateResponseCategory(category)
			responseItem := CreateResponseItem(item, responseCategory)
			responseDetails = append(responseDetails, CreateResponseOrderDetail(v, responseItem))
		}
		responseTable := CreateResponseTable(table)
		responseOrder := CreateResponseOrder(order, responseTable, responseDetails)
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func GetSpecificOrder(c *fiber.Ctx) error {
	return nil
}

/*
func GetActiveOrder(c *fiber.Ctx) error {
	tableRefer, err := c.ParamsInt("table_refer")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :table_refer is an integer")
	}
	var order models.Order
}
*/
