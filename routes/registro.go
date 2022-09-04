package routes

import (
	"encoding/json"
	"github.com/brayanzuritadev/daily/db"
	"github.com/brayanzuritadev/daily/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.Usuario

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "se necesita el email", 400)
		return
	}

	if len(u.Password) < 6 {
		http.Error(w, "su contraseña debe ser mayor a 6 digitos", 400)
		return
	}

	_, checkUser, status := db.ReviewUser(u)
	if checkUser == true {
		http.Error(w, status, 400)
		return
	}

	check, err, status := db.RegisterUser(u)
	if err != nil {
		http.Error(w, status+err.Error(), 400)
	}

	if check == false {
		http.Error(w, status, 400)
	}
	w.WriteHeader(http.StatusCreated)
}
