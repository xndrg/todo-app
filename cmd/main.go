package main

import (
	"log"

	todo "github.com/xndrg/crud-app"
	"github.com/xndrg/crud-app/internal/handler"
)

func main() {
	handler := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("4444", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
