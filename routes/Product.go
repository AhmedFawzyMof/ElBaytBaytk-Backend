package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Search struct {
	Search string `json:"search"`
}

func GetProduct(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	id, err := strconv.Atoi(req.PathValue("id"))

	if err != nil {
		er := errors.New("invalid id")
		middleware.SendError(er, res)
		return
	}

	db := database.Connect()
	defer db.Close()

	product := models.Product{}
	Product, err := product.ProductById(db, id)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.NewEncoder(res).Encode(Product); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func AllProducts(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	limit, err := strconv.Atoi(req.PathValue("limit"))

	if err != nil {
		er := errors.New("invalid id")
		middleware.SendError(er, res)
		return
	}

	db := database.Connect()
	defer db.Close()

	product := models.Product{}
	Products, err := product.GetAllProduct(db, limit)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	category := models.Category{}
	Categories, err := category.GetAllCategories(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["Products"] = Products
	Response["Categories"] = Categories

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func SearchProducts(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	limit, err := strconv.Atoi(req.PathValue("limit"))

	if err != nil {
		er := errors.New("invalid limit")
		middleware.SendError(er, res)
		return
	}

	search := Search{}
		
	if err := json.NewDecoder(req.Body).Decode(&search); err != nil {
		middleware.SendError(err, res)
		return
	}

	db := database.Connect()
	defer db.Close()

	product := models.Product{}
	Products, err := product.ProductsBySearch(db, search.Search, limit)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	if err := json.NewEncoder(res).Encode(Products); err != nil {
		middleware.SendError(err, res)
		return
	}
}