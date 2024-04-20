package models

import (
	"database/sql"
	"fmt"
)

type Carousel struct {
	Id      int           `json:"id"`
	Img     string        `json:"image"`
	Product sql.NullInt64 `json:"product"`
}

func (c Carousel) GetHomeCarousel(db *sql.DB) ([]Carousel, error) {
	carousel := []Carousel{}

	rows, err := db.Query("SELECT * FROM Carousel")

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error while prossing Carousel images")
	}

	defer rows.Close()

	for rows.Next() {
		var carousele Carousel

		if err := rows.Scan(&carousele.Id, &carousele.Product, &carousele.Img); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error while prossing Carousel images")
		}

		carousele.Img = "http://localhost:5500/assets" + carousele.Img
		carousel = append(carousel, carousele)
	}

	return carousel, nil
}
