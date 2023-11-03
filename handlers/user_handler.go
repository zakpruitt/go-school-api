package handlers

import (
	"school-api/models"
	"school-api/services"
)

// UserHandler uses the CrudHandler with the User type specified.
type UserHandler struct {
	*BaseCrudHandler[models.User]
}

// NewUserHandler creates a new handler for user operations.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		BaseCrudHandler: NewCrudHandler[models.User](userService),
	}
}
