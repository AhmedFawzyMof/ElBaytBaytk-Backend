package middleware

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func SendResponse(response chan []byte, Response map[string]interface{}) {
	res, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	response <- res
}

var SampleSecretKey = []byte("Ahmedfawzi made this website")

func GenerateJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 8760 * 2).Unix(),
	})
	tokenString, err := token.SignedString(SampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(AuthorizationToken string) (string, error) {

	tokenString := strings.Split(AuthorizationToken, "Token ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SampleSecretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := getEmailFromClaims(claims)
		return id, nil
	} else {
		return "", fmt.Errorf("invalid jwt token")
	}
}

func getEmailFromClaims(claims jwt.MapClaims) string {
	IdValue := claims["id"]

	id := IdValue.(string)

	return id
}
