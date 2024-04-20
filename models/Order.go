package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	Order        string  `json:"order"`
	Product      int     `json:"product"`
	Quantity     int     `json:"quantity"`
	Color        string  `json:"color"`
	ProductName  string  `json:"name"`
	ProductPrice float64 `json:"price"`
	ProductSlug  string  `json:"slug"`
}

type Order struct {
	Id          string      `json:"id"`
	User        string      `json:"user"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Method      string      `json:"method"`
	Phone       string      `json:"phone"`
	SparePhone  string      `json:"spare_phone"`
	Address     string      `json:"address"`
	IsPaid      bool        `json:"ispaid"`
	IsDelivered bool        `json:"isdelivered"`
	OrderStatus string      `json:"order_status"`
	OrderItems  []OrderItem `json:"items"`
	CreatedAt   time.Time   `json:"created_at"`
}

type OrdersS []Order

func (o Order) Create(db *sql.DB) (string, error) {

	o.Id = uuid.New().String()

	_, err := db.Exec("INSERT INTO Orders(id, user, name, method, phone, spare_phone, address, isPaid, isDelivered, order_status, created_at) VALUES(?,?,?,?,?,?,?,?,?,?,?)", o.Id, o.User, o.Name, o.Method, o.Phone, o.SparePhone, o.Address, o.IsPaid, o.IsDelivered, "pending", o.CreatedAt)

	if err != nil {
		return "", fmt.Errorf("error occurred while processing the order please try again or contact us")
	}

	items, err := db.Prepare("INSERT INTO OrdersItem(`order`, product, quantity, color) VALUES(?, ?, ?, ?)")

	if err != nil {
		return "", fmt.Errorf("error occurred while processing the order please try again or contact us")
	}

	for _, item := range o.OrderItems {
		_, err := items.Exec(o.Id, item.Product, item.Quantity, item.Color)

		if err != nil {
			return "", fmt.Errorf("error occurred while processing the order please try again or contact us")
		}
	}

	id := strings.Split(o.Id, "-")[0]

	return id, nil
}

func (o Order) GetHistory(db *sql.DB) ([]Order, error) {

	var Orders []Order

	orders, err := db.Query("SELECT * FROM Orders WHERE Orders.user = ?", o.User)
	if err != nil {
		return nil, fmt.Errorf("error occurred while processing the order please try again or contact us")
	}

	defer orders.Close()

	for orders.Next() {
		var Order Order

		if err := orders.Scan(&Order.Id, &Order.User, &Order.Name, &Order.Method, &Order.Phone, &Order.SparePhone, &Order.Address, &Order.IsPaid, &Order.IsDelivered, &Order.OrderStatus, &Order.CreatedAt); err != nil {
			return nil, fmt.Errorf("error occurred while processing the order please try again or contact us")
		}

		Orders = append(Orders, Order)
	}

	return Orders, nil
}

func (o *OrdersS) GetOrderProducts(db *sql.DB) error {

	productsPre, err := db.Prepare("SELECT OrdersItem.`order`, OrdersItem.quantity, Products.name, Products.price FROM OrdersItem INNER JOIN Products ON OrdersItem.product = Products.id WHERE OrdersItem.`order` = ?")

	if err != nil {
		return fmt.Errorf("error occurred while processing the order please try again or contact us")
	}
	for i := range *o {

		products, err := productsPre.Query((*o)[i].Id)

		if err != nil {
			return fmt.Errorf("error occurred while processing the order please try again or contact us")
		}

		defer products.Close()

		for products.Next() {
			var OrderItem OrderItem

			if err := products.Scan(&OrderItem.Order, &OrderItem.Quantity, &OrderItem.ProductName, &OrderItem.ProductPrice); err != nil {
				return fmt.Errorf("error occurred while processing the order please try again or contact us")
			}
			(*o)[i].OrderItems = append((*o)[i].OrderItems, OrderItem)
		}
	}
	return nil
}
