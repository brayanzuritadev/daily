package handlers

import (
	"github.com/brayanzuritadev/daily/middlew"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/brayanzuritadev/daily/routes"
)

func Handlers() {

	//esta es una manera de conectar
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckConnection(routes.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
