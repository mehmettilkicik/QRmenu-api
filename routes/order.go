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

// Create Order
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

//Create Order End

// Find All Orders
func findOrders(orders *[]models.Order) error {
	config.Database.Db.Find(&orders)
	if len(*orders) == 0 {
		return errors.New("there is no order")
	}
	return nil
}

// Find All Orders End

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

// Find Active Or Inactive Orders
func FindActiveOrInactiveOrders(orders *[]models.Order, isPaid bool) error {
	config.Database.Db.Find(&orders, "is_paid=?", isPaid)
	if len(*orders) == 0 {
		if !isPaid {
			return errors.New("no active orders")
		} else if isPaid {
			return errors.New("no inactive orders")
		}
	}
	return nil
}

//Find Active Or Inactive Orders End

// Get Active or Inactive Orders

func GetActiveOrInactiveOrders(c *fiber.Ctx) error {
	aori, err := c.ParamsInt("is_paid")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :is_paid is an integer")
	}
	var acorin bool

	if aori == 1 {
		acorin = true
	} else if aori == 0 {
		acorin = false
	}

	var orders []models.Order
	if err := FindActiveOrInactiveOrders(&orders, acorin); err != nil {
		if err.Error() == "no active orders" {
			return c.Status(404).JSON("no active order found")
		} else if err.Error() == "no inactive orders" {
			return c.Status(404).JSON("no inactive order found")
		}
		return c.Status(500).JSON("internal server error")
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

//Get Active Or Inactive Orders End

// Find Order By Table
func findOrderByTable(tableRefer int, orders *[]models.Order, isPaid bool) error {
	config.Database.Db.Where("is_paid=?", isPaid).Find(&orders, "table_refer", tableRefer)
	if len(*orders) == 0 {
		if !isPaid {
			return errors.New("no active orders in this table")
		} else if isPaid {
			return errors.New("no inactive orders in this table")
		}
	}
	return nil
}

//Find Order By Table End

//Get Orders By Table

func GetOrdersByTable(c *fiber.Ctx) error {
	aori, err := c.ParamsInt("is_paid")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :is_paid is an integer")
	}
	var acorin bool

	if aori == 1 {
		acorin = true
	} else if aori == 0 {
		acorin = false
	}

	tableRefer, err := c.ParamsInt("table_refer")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :table_refer is an integer")
	}
	var orders []models.Order
	if err := findOrderByTable(tableRefer, &orders, acorin); err != nil {
		if err.Error() == "no active orders" {
			return c.Status(404).JSON("no active order found in this table")
		} else if err.Error() == "no inactive orders" {
			return c.Status(404).JSON("no inactive order found in this table")
		}
		return c.Status(500).JSON("internal server error")
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

//Get Orders By Table End

//Find Specific Order

func findSpecificOrder(order *models.Order, tableRefer int, isPaid bool, id int) error {
	config.Database.Db.Where("is_paid=?", isPaid).Where("table_refer=?", tableRefer).Find(&order, "id=?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

//Find Specific Order End

// Get Specific Order

func GetSpecificOrder(c *fiber.Ctx) error {
	aori, err := c.ParamsInt("is_paid")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :is_paid is an integer")
	}
	var acorin bool

	if aori == 1 {
		acorin = true
	} else if aori == 0 {
		acorin = false
	}

	tableRefer, err := c.ParamsInt("table_refer")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :table_refer is an integer")
	}

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var order models.Order

	if err := findSpecificOrder(&order, tableRefer, acorin, id); err != nil {
		if err.Error() == ("order does not exist") {
			return c.Status(404).JSON("there is no such an order")
		}
		return c.Status(500).JSON("internal server error")
	}

	responseDetails := []OrderDetail{}
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

	return c.Status(200).JSON(responseOrder)
}

// Get Specific Order End
