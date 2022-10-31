package main

import (
	"Skipper_cms_users/pkg/handlers"
	"Skipper_cms_users/pkg/models"
	"Skipper_cms_users/pkg/repository"
	"Skipper_cms_users/pkg/servises"
)

// @title Users service
// @version 1.0
// @description Methods for control users for skipper cms
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := models.GetDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlerses := handlers.NewHandler(services)
	handlerses.InitRoutes()
}
