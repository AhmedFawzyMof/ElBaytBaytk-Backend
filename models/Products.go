package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type FilterData struct {
	Min_price int    `json:"min_price"`
	Max_price int    `json:"max_price"`
	Category  string `json:"category"`
}

type Product struct {
	Id            int            `json:"id"`
	SubCategory   int            `json:"subcategory"`
	Category      int            `json:"category"`
	Name          string         `json:"name"`
	NameAr        string         `json:"nameAr"`
	Description   string         `json:"description"`
	DescriptionAr string         `json:"descriptionAr"`
	Price         float64        `json:"price"`
	Discount      float64        `json:"discount"`
	Image         string         `json:"image"`
	Warranty      sql.NullString `json:"warranty"`
	Brand         string         `json:"brand"`
	Material      string         `json:"material"`
	Color         sql.NullString `json:"color"`
}

func (p Product) GetAllProduct(db *sql.DB, limit int) ([]Product, error) {
	var Products []Product
	var oldLimit int = 0

	if limit > 20 {
		oldLimit = limit / 2
	}

	const stableLimit = 20

	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.category, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.discount > 0  GROUP BY Products.id ORDER BY Products.discount DESC LIMIT ?,?")

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error while prossing products")
	}

	products, err := productsPre.Query(oldLimit, stableLimit)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error while prossing products")
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Id, &Product.Name, &Product.Category, &Product.NameAr, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {

			fmt.Println(err.Error())
			return nil, errors.New("error while prossing products")
		}

		Product.Image = "http://localhost:5500/assets" + Product.Image
		Products = append(Products, Product)
	}

	return Products, nil

}

func (p Product) ProductsByCategorys(db *sql.DB, id, limit int) ([]Product, error) {
	var Products []Product
	var oldLimit int = 0

	if limit > 20 {
		oldLimit = limit / 2
	}

	const stableLimit = 20

	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.category = ? GROUP BY Products.id LIMIT ?,?")

	if err != nil {
		return nil, errors.New("error while prossing products")
	}

	products, err := productsPre.Query(id, oldLimit, stableLimit)

	if err != nil {
		return nil, errors.New("error while prossing products")
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
			return nil, errors.New("error while prossing products")
		}

		Product.Image = "http://localhost:5500/assets" + Product.Image
		Products = append(Products, Product)
	}

	return Products, nil
}

func (p Product) ProductsBySubCategorys(db *sql.DB, id, limit int) ([]Product, error) {
	var Products []Product
	var oldLimit int = 0

	if limit > 20 {
		oldLimit = limit / 2
	}

	const stableLimit = 20

	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.subcategory = ? GROUP BY Products.id LIMIT ?,?")

	if err != nil {
		return nil, errors.New("error while prossing products")
	}

	products, err := productsPre.Query(id, oldLimit, stableLimit)

	if err != nil {
		return nil, errors.New("error while prossing products")
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
			return nil, errors.New("error while prossing products")
		}

		Product.Image = "http://localhost:5500/assets" + Product.Image
		Products = append(Products, Product)
	}

	return Products, nil
}

func (p Product) ProductById(db *sql.DB, id int) (Product, error) {
	var product Product

	preprow, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, Products.warranty, Products.brand, Products.material, ProductImages.image, ProductImages.color FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.id = ?")

	if err != nil {
		return Product{}, errors.New("error while prossing product")
	}

	row := preprow.QueryRow(id)

	if err := row.Scan(&product.Id, &product.Name, &product.NameAr, &product.Description, &product.DescriptionAr, &product.Price, &product.Discount, &product.Warranty, &product.Brand, &product.Material, &product.Image, &product.Color); err != nil {
		fmt.Println(err.Error())
		return Product{}, errors.New("error while prossing product")
	}

	product.Image = "http://localhost:5500/assets" + product.Image

	return product, nil
}

func (p Product) ProductByOffer(db *sql.DB, subcategory, limit int) ([]Product, error) {
	var Products []Product
	var oldLimit int = 0

	if limit > 20 {
		oldLimit = limit / 2
	}

	const stableLimit = 20

	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.subcategory = ? AND Products.discount > 0  GROUP BY Products.id ORDER BY Products.discount DESC LIMIT ?,?")

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error while prossing products")
	}

	products, err := productsPre.Query(subcategory, oldLimit, stableLimit)

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

		Product.Image = "http://localhost:5500/assets" + Product.Image
		Products = append(Products, Product)
	}

	return Products, nil
}

func (p Product) ProductsBySearch(db *sql.DB, search string, limit int) ([]Product, error) {
	var Products []Product
	var oldLimit int = 0

	if limit > 20 {
		oldLimit = limit / 2
	}

	const stableLimit = 20
	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.name LIKE ? OR Products.description LIKE ? GROUP BY Products.id LIMIT ?,?")

	if err != nil {
		return nil, fmt.Errorf("error while prossing products")
	}

	products, err := productsPre.Query("%"+search+"%", "%"+search+"%", oldLimit, stableLimit)

	if err != nil {
		return nil, fmt.Errorf("error while prossing products")
	}

	defer products.Close()

	for products.Next() {
		var Product Product

		if err := products.Scan(&Product.Id, &Product.Name, &Product.NameAr, &Product.Description, &Product.DescriptionAr, &Product.Price, &Product.Discount, &Product.Image); err != nil {
			return nil, fmt.Errorf("error while prossing products")
		}

		Product.Image = "http://localhost:5500/assets" + Product.Image

		Products = append(Products, Product)
	}

	return Products, nil
}
