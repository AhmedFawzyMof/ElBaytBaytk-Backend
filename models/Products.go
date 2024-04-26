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

	productsPre, err := db.Prepare("SELECT Products.id, Products.name, Products.category, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, ProductImages.image FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product GROUP BY Products.id ORDER BY Products.discount DESC LIMIT ?,?")

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

		Product.Image = "https://elbaytbaytk-backend.onrender.com/assets" + Product.Image
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

		Product.Image = "https://elbaytbaytk-backend.onrender.com/assets" + Product.Image
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

		Product.Image = "https://elbaytbaytk-backend.onrender.com/assets" + Product.Image
		Products = append(Products, Product)
	}

	return Products, nil
}

func (p Product) ProductById(db *sql.DB, id int) (map[int]map[string]interface{}, error) {
	ProductMap := make(map[int]map[string]interface{})

	preprow, err := db.Prepare("SELECT Products.id, Products.name, Products.nameAr, Products.description, Products.descriptionAr, Products.price, Products.discount, Products.warranty, Products.brand, Products.material, ProductImages.image, ProductImages.color FROM Products INNER JOIN ProductImages ON Products.id = ProductImages.product WHERE Products.id = ?")

	if err != nil {
		return nil, errors.New("error while prossing product")
	}

	rows, err := preprow.Query(id)

	if err != nil {
		return nil, errors.New("error while prossing product")
	}

	defer rows.Close()

	for rows.Next() {
		var product Product

		if err := rows.Scan(&product.Id, &product.Name, &product.NameAr, &product.Description, &product.DescriptionAr, &product.Price, &product.Discount, &product.Warranty, &product.Brand, &product.Material, &product.Image, &product.Color); err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("error while prossing product")
		}

		product.Image = "https://elbaytbaytk-backend.onrender.com/assets" + product.Image

		if _, ok := ProductMap[product.Id]; !ok {
			ProductMap[product.Id] = map[string]interface{}{
				"id":            product.Id,
				"name":          product.Name,
				"nameAr":        product.NameAr,
				"description":   product.Description,
				"descriptionAr": product.DescriptionAr,
				"price":         product.Price,
				"discount":      product.Discount,
				"warranty":      product.Warranty,
				"brand":         product.Brand,
				"material":      product.Material,
				"images":        []map[string]string{},
				"colors":        []sql.NullString{},
			}
		}
		ProductMap[product.Id]["images"] = append(ProductMap[product.Id]["images"].([]map[string]string), map[string]string{"image": product.Image})
		ProductMap[product.Id]["colors"] = append(ProductMap[product.Id]["colors"].([]sql.NullString), product.Color)
	}

	return ProductMap, nil
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

		Product.Image = "https://elbaytbaytk-backend.onrender.com/assets" + Product.Image
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

		Product.Image = "https://elbaytbaytk-backend.onrender.com/assets" + Product.Image

		Products = append(Products, Product)
	}

	return Products, nil
}
