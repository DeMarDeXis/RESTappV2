package main

import (
	"log"

	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/DeMarDeXis/RESTV1/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(gorestapiv2.Server)
	if err := srv.Start("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
