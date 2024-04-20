package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type Contact struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func (c Contact) AddContact(db *sql.DB) []byte {
	Response := make(map[string]interface{})

	_, err := db.Exec("INSERT INTO ContactUs(name, email,phone,message) VALUES(?,?,?,?)", c.Name, c.Email, c.Phone, c.Message)

	if err != nil {
		fmt.Println(err.Error())
	}

	Response["Done"] = true

	responseBytes, err := json.Marshal(Response)

	if err != nil {
		fmt.Println(err.Error())
	}

	return responseBytes
}
