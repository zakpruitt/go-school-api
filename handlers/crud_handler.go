package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"school-api/services"
	"strconv"
)

// BaseCrudHandler is a generic handler for CRUD operations.
type BaseCrudHandler[T any] struct {
	CrudServiceInterface services.CrudServiceInterface[T]
}

// NewCrudHandler creates a new BaseCrudHandler with the provided CrudServiceInterface.
func NewCrudHandler[T any](crudService services.CrudServiceInterface[T]) *BaseCrudHandler[T] {
	return &BaseCrudHandler[T]{
		CrudServiceInterface: crudService,
	}
}

// GetOneObjectById handles GET requests to retrieve an object by its ID.
func (h *BaseCrudHandler[T]) GetOneObjectById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		JSONErrorResponse(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	obj, err := h.CrudServiceInterface.GetByID(id)
	if err != nil {
		JSONErrorResponse(w, "Object not found", http.StatusNotFound)
		return
	}

	JSONResponse(w, obj, http.StatusOK)
}
