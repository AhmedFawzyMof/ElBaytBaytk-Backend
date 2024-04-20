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


func SubCategoryByid(res http.ResponseWriter, req *http.Request) {
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

	subcategory := models.SubCategory{}
	SubCategory, err := subcategory.GetSubCategoryById(db, id)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	product := models.Product{}
	Products, err := product.ProductsBySubCategorys(db, id, limit)
	
	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["SubCategory"] = SubCategory
	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}