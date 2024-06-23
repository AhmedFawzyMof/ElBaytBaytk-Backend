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

	for i, item := range o.OrderItems {
		fmt.Printf("Processing item %d: %+v\n", i, item)
		_, err := items.Exec(o.Id, item.Product, item.Quantity, item.Color)

		if err != nil {
			fmt.Printf("Error processing item %d: %v\n", i, err)
			return "", fmt.Errorf("error occurred while processing the order please try again or contact us")
		}
	}

	id := strings.Split(o.Id, "-")[0]

	return id, nil
}

func (o Order) GetHistory(db *sql.DB) ([]Order, error) {

	var Orders []Order

	stmt, err := db.Prepare("SELECT Orders.id, Orders.name, Orders.address, Orders.phone, Orders.spare_phone, Orders.method, Orders.order_status, Orders.isPaid, Orders.isDelivered, Orders.created_at, OrdersItem.color, OrdersItem.quantity, Products.name, Products.price FROM Orders INNER JOIN OrdersItem ON OrdersItem.`order` = Orders.id INNER JOIN Products ON Products.id = OrdersItem.product WHERE Orders.user = ?")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(o.User)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Order Order
		var Item OrderItem

		if err := rows.Scan(&Order.Id, &Order.Name, &Order.Address, &Order.Phone, &Order.SparePhone, &Order.Method, &Order.OrderStatus, &Order.IsPaid, &Order.IsDelivered, &Order.CreatedAt, &Item.Color, &Item.Quantity, &Item.ProductName, &Item.ProductPrice); err != nil {
			fmt.Println(err.Error())
			return nil, fmt.Errorf("error occurred while processing the order please try again or contact us")
		}

		index, found := FindOrderById(Orders, Order.Id)

		if found {
			Orders[index].OrderItems = append(Orders[index].OrderItems, Item)
		} else {
			Order.OrderItems = append(Order.OrderItems, Item)
			Orders = append(Orders, Order)
		}
	}

	return Orders, nil
}

func FindOrderById(orders []Order, id string) (int, bool) {
	for i, order := range orders {
		if order.Id == id {
			return i, true
		}
	}
	return -1, false
}
