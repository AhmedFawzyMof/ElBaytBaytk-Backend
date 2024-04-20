package models

import (
	"database/sql"
	"fmt"
)

type Category struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	NameAr string `json:"nameAr"`
}

func (c Category) GetAllCategories(db *sql.DB) ([]Category, error) {
	var Categories []Category

	categories, err := db.Query("SELECT * FROM Categories")

	if err != nil {
		return nil, fmt.Errorf("error while prossing categories")
	}

	defer categories.Close()

	for categories.Next() {
		var Category Category

		if err := categories.Scan(&Category.Id, &Category.Name, &Category.NameAr); err != nil {
			return nil, fmt.Errorf("error while prossing categories")
		}

		Categories = append(Categories, Category)
	}

	return Categories, nil
}

func (c Category) GetCategoryById(db *sql.DB, categoryId int) (Category, error) {

	categories := db.QueryRow("SELECT * FROM Categories WHERE id = ?", categoryId)

	var category Category

	if err := categories.Scan(&category.Id, &category.Name, &category.NameAr); err != nil {
		fmt.Println(err.Error())
		return Category{}, fmt.Errorf("error while prossing categories")
	}

	return category, nil
}
