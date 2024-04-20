package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
	"time"
)

func Order(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	Authorization := req.Header.Get("Authorization")

	orderReq := models.Order{}

	if err := json.NewDecoder(req.Body).Decode(&orderReq); err != nil {
		middleware.SendError(err, res)
		return
	}

	if Authorization != "" {
		id, err := middleware.VerifyToken(Authorization)

		if err != nil {
			middleware.SendError(err, res)
			return
		}

		orderReq.User = id
	}

	if orderReq.Method != "cash" {
		orderReq.IsPaid = true
	}

	orderReq.CreatedAt = time.Now()

	orderId, err := orderReq.Create(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.NewEncoder(res).Encode(map[string]interface{}{
		"success": true,
		"id":      orderId,
	}); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func OrderHistory(res http.ResponseWriter, req *http.Request) {
	Authorization := req.Header.Get("Authorization")
	if Authorization == "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}

	user, err := middleware.VerifyToken(Authorization)

	if err != nil {
		res.WriteHeader(http.StatusForbidden)
		return
	}

	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	order := models.Order{}
	order.User = user
	Orders, err := order.GetHistory(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Order := models.OrdersS(Orders)

	if err := Order.GetOrderProducts(db); err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := map[string]interface{}{
		"Orders": Order,
	}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}
