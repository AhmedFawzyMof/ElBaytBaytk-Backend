package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type SubCategory struct {
	Id             int    `json:"id"`
	CategoryId     int    `json:"category_id"`
	CategoryName   string `json:"category_name"`
	CategoryNameAr string `json:"category_name_ar"`
	Name           string `json:"name"`
	NameAr         string `json:"name_ar"`
	Img            string `json:"image"`
}

func (s SubCategory) GetHomeSubCategory(db *sql.DB) ([]SubCategory, error) {
	SubCategories := []SubCategory{}

	rows, err := db.Query("SELECT sc.id, sc.name, sc.nameAr, sc.category, c.name AS categoryName, c.nameAr AS categoryNameAr, sc.img FROM ( SELECT *, ROW_NUMBER() OVER(PARTITION BY category ORDER BY id) as rn FROM SubCategory ) AS sc JOIN Categories c ON sc.category = c.id WHERE sc.rn <= 4")

	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("error while prossing subcategories")
	}

	defer rows.Close()

	for rows.Next() {
		var SubCategory SubCategory
		if err := rows.Scan(&SubCategory.Id, &SubCategory.Name, &SubCategory.NameAr, &SubCategory.CategoryId, &SubCategory.CategoryName, &SubCategory.CategoryNameAr, &SubCategory.Img); err != nil {
			fmt.Println(err.Error())
			return nil, fmt.Errorf("error while prossing subcategories")
		}
		SubCategory.Img = "http://localhost:5500/assets" + SubCategory.Img
		SubCategories = append(SubCategories, SubCategory)
	}

	return SubCategories, nil
}

func (s SubCategory) GetAllSubCategory(db *sql.DB) ([]SubCategory, error) {
	SubCategories := []SubCategory{}
	rows, err := db.Query("SELECT sc.id, sc.name, sc.nameAr, sc.category, c.name AS categoryName, c.nameAr AS categoryNameAr, sc.img FROM ( SELECT *, ROW_NUMBER() OVER(PARTITION BY category ORDER BY id) as rn FROM SubCategory ) AS sc JOIN Categories c ON sc.category = c.id ORDER BY sc.id")

	if err != nil {
		return nil, fmt.Errorf("error while prossing subcategories")
	}

	defer rows.Close()

	for rows.Next() {
		var SubCategory SubCategory
		if err := rows.Scan(&SubCategory.Id, &SubCategory.Name, &SubCategory.NameAr, &SubCategory.CategoryId, &SubCategory.CategoryName, &SubCategory.CategoryNameAr, &SubCategory.Img); err != nil {
			return nil, fmt.Errorf("error while prossing subcategories")
		}
		SubCategory.Img = "http://localhost:5500/assets" + SubCategory.Img
		SubCategories = append(SubCategories, SubCategory)
	}

	return SubCategories, nil
}

func (s SubCategory) GetSubCategoryByCategory(db *sql.DB, categoryId int) ([]SubCategory, error) {
	var SubCategories []SubCategory

	rows, err := db.Query("SELECT * FROM SubCategory WHERE category = ?", categoryId)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error while prossing subcategories")
	}

	defer rows.Close()

	for rows.Next() {
		var SubCategory SubCategory

		if err := rows.Scan(&SubCategory.Id, &SubCategory.Name, &SubCategory.NameAr, &SubCategory.CategoryId, &SubCategory.Img); err != nil {
			return nil, errors.New("error while prossing subcategories")
		}
		SubCategory.Img = "http://localhost:5500/assets" + SubCategory.Img
		SubCategories = append(SubCategories, SubCategory)
	}

	return SubCategories, nil

}

func (s SubCategory) GetSubCategoryById(db *sql.DB, id int) (SubCategory, error) {
	var subCategory SubCategory
	row := db.QueryRow("SELECT * FROM SubCategory WHERE id = ?", id)

	if err := row.Scan(&subCategory.Id, &subCategory.Name, &subCategory.NameAr, &subCategory.CategoryId, &subCategory.Img); err != nil {
		return SubCategory{}, errors.New("error while prossing subcategories")
	}

	subCategory.Img = "http://localhost:5500/assets" + subCategory.Img

	return subCategory, nil
}
