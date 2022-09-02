package main

import (
	"log"

	"github.com/brayanzuritadev/daily/db"
	"github.com/brayanzuritadev/daily/handlers"
)

func main() {
	if db.CheckConnection()==0 {
		log.Fatal("sin conexion a la bd")
		return
	}
	handlers.Handlers()

}
