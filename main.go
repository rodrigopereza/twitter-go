package main

import (
	"log"

	db "github.com/rodrigopereza/twitter-go/bd"
	"github.com/rodrigopereza/twitter-go/handlers"
)

func main() {
	if db.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
