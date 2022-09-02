package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/brayanzuritadev/daily/middlew"
	"github.com/brayanzuritadev/daily/routes"
)

func Handlers() {

	//esta es una manera de conectar
	router := mux.NewRouter()

	router.HandleFunc("register", middlew.CheckConnection(routes.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
	
}
