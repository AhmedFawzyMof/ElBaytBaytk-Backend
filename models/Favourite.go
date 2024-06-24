package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Favourite struct {
	Id      int    `json:"id"`
	Product int    `json:"product"`
	User    string `json:"user"`
}

func (f Favourite) GetAllFavourite(db *sql.DB) ([]Product, error) {
	var Products []Product

	products, err := db.Query("SELECT Favourite.product, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Favourite INNER JOIN Products ON Favourite.product = Products.id INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Favourite.user = ? GROUP BY Products.id", f.User)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error while prossing products")
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("error while prossing products")
		}

		Product.Image = "https://elbaytbaytk-backend.onrender.com/assets" + Product.Image
		Products = append(Products, Product)
	}

	return Products, nil
}

func (f Favourite) AddToFavourite(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO Favourite (user, product) VALUES (?, ?)", f.User, f.Product)

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("error while adding product to favourite")
	}

	return nil
}

func (f Favourite) FindInFavourite(db *sql.DB) (bool, error) {

	row := db.QueryRow("SELECT product FROM Favourite WHERE (user, product) = (?, ?)", f.User, f.Product)

	Favourite := Favourite{}

	if err := row.Scan(&Favourite.Product); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, fmt.Errorf("error while finding product in favourite")
		}
	}

	return true, nil
}
