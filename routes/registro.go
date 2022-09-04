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

	_, found, _ := db.ReviewUser(u)
	if found == true {
		http.Error(w, "El email del usuario ya existe dentro la base de datos", 400)
		return
	}

	_, status, err := db.RegisterUser(u)
	if err != nil {
		http.Error(w, "no se pudo registrar al usuario dentro la base de datos "+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "no se a logrado insertar el registro del usuario", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
