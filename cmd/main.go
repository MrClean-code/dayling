package main

import (
	"github.com/MrClean-code/dayling"
	"github.com/MrClean-code/dayling/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(dayling.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured white running http server: %s", err.Error())
	}

}
