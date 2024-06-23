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

func Category(res http.ResponseWriter, req *http.Request) {
	db := database.Connect()
	defer db.Close()

	subcategory := models.SubCategory{}
	SubCategories, err := subcategory.GetAllSubCategory(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["SubCategories"] = SubCategories

	res.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func Categories(res http.ResponseWriter, req *http.Request) {
	db := database.Connect()
	defer db.Close()

	category := models.Category{}
	Categories, err := category.GetAllCategories(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["Categories"] = Categories

	res.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func CategoryPage(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	id, err := strconv.Atoi(req.PathValue("id"))

	if err != nil {
		er := errors.New("invalid id")
		middleware.SendError(er, res)
		return
	}

	limit, err := strconv.Atoi(req.PathValue("limit"))

	if err != nil {
		er := errors.New("invalid id")
		middleware.SendError(er, res)
		return
	}

	db := database.Connect()
	defer db.Close()

	category := models.Category{}
	Category, err := category.GetCategoryById(db, id)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	products := models.Product{}
	Products, err := products.ProductsByCategorys(db, id, limit)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	subcategories := models.SubCategory{}
	SubCategories, err := subcategories.GetSubCategoryByCategory(db, id)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["Category"] = Category
	Response["Products"] = Products
	Response["SubCategories"] = SubCategories

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}

}
