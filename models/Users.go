package models

import (
	"HAstore/middleware"
	"crypto/sha256"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Users struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u Users) Register(db *sql.DB) (string, error) {
	sha := sha256.New()
	sha.Write([]byte(u.Password))
	pass := sha.Sum(nil)
	u.Password = fmt.Sprintf("%x", pass)
	u.Id = uuid.New().String()

	_, err := db.Exec("INSERT INTO Users(id, username, email, password) VALUES(?,?,?,?)", u.Id, u.Username, u.Email, u.Password)

	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("can't register with provided credentials")
	}

	token, err := middleware.GenerateJWT(u.Id)

	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("can't register with provided credentials")
	}

	return token, nil
}

func (u Users) Login(db *sql.DB) (string, error) {
	sha := sha256.New()
	sha.Write([]byte(u.Password))
	pass := sha.Sum(nil)
	u.Password = fmt.Sprintf("%x", pass)

	userQuery := db.QueryRow("SELECT email,id FROM Users WHERE (email, password) = (?,?)", u.Email, u.Password)

	var User Users

	if err := userQuery.Scan(&User.Email,&User.Id); err != nil {
		return "", fmt.Errorf("can't login with provided credentials")
	}

	token, err := middleware.GenerateJWT(User.Id)

	if err != nil {
		return "", fmt.Errorf("can't login with provided credentials")
	}

	return token, nil
}
