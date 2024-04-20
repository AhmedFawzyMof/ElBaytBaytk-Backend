package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	subcategory := models.SubCategory{}
	SubCategories, err := subcategory.GetHomeSubCategory(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	offers := models.Offers{}
	Offers, err := offers.GetAllOffers(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	carousel := models.Carousel{}
	Carousels, err := carousel.GetHomeCarousel(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := make(map[string]interface{})
	Response["SubCategories"] = SubCategories
	Response["Offers"] = Offers
	Response["Carousels"] = Carousels

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}
