package main

import (
	"github.com/MrClean-code/dayling"
	"github.com/MrClean-code/dayling/pkg/handler"
	"github.com/MrClean-code/dayling/pkg/repository"
	"github.com/MrClean-code/dayling/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(dayling.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured white running http server: %s", err.Error())
	}

}
