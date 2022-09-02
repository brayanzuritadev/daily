package middlew

import (
	"net/http"

	"github.com/brayanzuritadev/daily/db"
)

func CheckConnection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if db.CheckConnection() == 0{
			http.Error(w, "error en conexion", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}