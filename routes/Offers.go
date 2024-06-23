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

func ProductByOffer(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	subcategory, err := strconv.Atoi(req.PathValue("subcategory"))

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

	product := models.Product{}

	Products, err := product.ProductByOffer(db, subcategory, limit)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := map[string]interface{} {
		"Products": Products,
	}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}

}
