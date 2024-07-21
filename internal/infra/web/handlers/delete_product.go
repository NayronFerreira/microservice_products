package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/NayronFerreira/microservice_products/internal/entity"
	"github.com/NayronFerreira/microservice_products/internal/usecase"
)

type DeleteProductHandler struct {
	repo entity.ProductRepoInterface
}

func NewDeleteProductHandler(repo entity.ProductRepoInterface) *DeleteProductHandler {
	return &DeleteProductHandler{repo: repo}
}

func (h DeleteProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	var inputDTO usecase.DeleteProductInputDTO
	if err := json.NewDecoder(r.Body).Decode(&inputDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if inputDTO.ID == "" {
		http.Error(w, errors.New("ID is required").Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteProduct(inputDTO.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h DeleteProductHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("DELETE /product", h.DeleteProductHandler)
}
