package routes

import (
	"HAstore/database"
	"HAstore/middleware"
	"HAstore/models"
	"encoding/json"
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	Authorization := req.Header.Get("Authorization")

	if Authorization != "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}
	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	user := models.Users{}

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		middleware.SendError(err, res)
		return
	}

	Token, err := user.Login(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := map[string]interface{}{"token": Token}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func Register(res http.ResponseWriter, req *http.Request) {
	Authorization := req.Header.Get("Authorization")

	if Authorization != "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}
	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	user := models.Users{}

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		middleware.SendError(err, res)
		return
	}

	Token, err := user.Register(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := map[string]interface{}{"token": Token}

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}
}

func Favourite(res http.ResponseWriter, req *http.Request) {
	Authorization := req.Header.Get("Authorization")

	if Authorization == "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}

	user, err := middleware.VerifyToken(Authorization)

	if err != nil {
		res.WriteHeader(http.StatusForbidden)
		return
	}

	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	favourite := models.Favourite{}

	// if err := json.NewDecoder(req.Body).Decode(&favourite); err != nil {
	// 	middleware.SendError(err, res)
	// 	return
	// }

	favourite.User = user

	Products, err := favourite.GetAllFavourite(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	Response := map[string]interface{}{}

	Response["Products"] = Products

	if err := json.NewEncoder(res).Encode(Response); err != nil {
		middleware.SendError(err, res)
		return
	}

	// if err := favourite.AddToFavourite(db); err != nil {
	// 	middleware.SendError(err, res)
	// 	return
	// }
}

func AddToFavourite(res http.ResponseWriter, req *http.Request) {
	Authorization := req.Header.Get("Authorization")

	if Authorization == "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}

	user, err := middleware.VerifyToken(Authorization)

	if err != nil {
		res.WriteHeader(http.StatusForbidden)
		return
	}

	res.WriteHeader(http.StatusOK)

	db := database.Connect()
	defer db.Close()

	favourite := models.Favourite{}

	favourite.User = user

	if err := json.NewDecoder(req.Body).Decode(&favourite); err != nil {
		middleware.SendError(err, res)
		return
	}

	found, err := favourite.FindInFavourite(db)

	if err != nil {
		middleware.SendError(err, res)
		return
	}

	if !found {
		if err := favourite.AddToFavourite(db); err != nil {
			middleware.SendError(err, res)
			return
		} else if err := json.NewEncoder(res).Encode(map[string]bool{"success": true}); err != nil {
			middleware.SendError(err, res)
			return
		}
		return
	}

	if err := json.NewEncoder(res).Encode(map[string]bool{"success": false}); err != nil {
		middleware.SendError(err, res)
		return
	}
}
