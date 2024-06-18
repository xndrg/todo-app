package main

import (
	"log"

	"github.com/spf13/viper"
	todo "github.com/xndrg/crud-app"
	"github.com/xndrg/crud-app/internal/handler"
	"github.com/xndrg/crud-app/internal/repository"
	"github.com/xndrg/crud-app/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
