package models

import (
	"database/sql"
	"fmt"
)

type Offers struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"image"`
	Subcategory int    `json:"subcategory"`
}

func (o Offers) GetAllOffers(db *sql.DB) ([]Offers, error) {

	var TOffers []Offers

	offers, err := db.Query("SELECT * FROM Offers")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error while prossing offers")
	}

	defer offers.Close()

	for offers.Next() {
		var offer Offers

		if err := offers.Scan(&offer.Id, &offer.Name, &offer.Subcategory, &offer.Img); err != nil {
			return nil, fmt.Errorf("error while prossing offers")
		}

		offer.Img = "http://localhost:5500/assets" + offer.Img

		TOffers = append(TOffers, offer)
	}

	return TOffers, nil
}
