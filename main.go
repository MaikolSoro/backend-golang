package main

import (
	"log"

	"github.com/maikol/backend-golang/bd"
	"github.com/maikol/backend-golang/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion ala BD")
		return
	}
	handlers.Manejadores()
}
