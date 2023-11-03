package config

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"school-api/handlers"
	"school-api/services"
)

func InitRouter(db *gorm.DB) *mux.Router {
	// Services
	userService := services.NewUserService(db)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/user/{id:[0-9]+}", userHandler.GetOneObjectById).Methods("GET")

	return router
}
