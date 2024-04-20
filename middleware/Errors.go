package middleware

import "net/http"

func SendError(err error, res http.ResponseWriter) {
	http.Error(res, err.Error(), http.StatusInternalServerError)
}
