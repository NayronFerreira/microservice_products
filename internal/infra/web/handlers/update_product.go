package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/microservice_products/internal/entity"
	"github.com/NayronFerreira/microservice_products/internal/usecase"
)

type UpdateProductHandler struct {
	productRepo entity.ProductRepoInterface
}

func NewUpdateProductHandler(productRepo entity.ProductRepoInterface) *UpdateProductHandler {
	return &UpdateProductHandler{
		productRepo: productRepo,
	}
}

func (h *UpdateProductHandler) UpdateProductHandle(w http.ResponseWriter, r *http.Request) {

	var inputDTO usecase.UpdateProductInputDTO
	if err := json.NewDecoder(r.Body).Decode(&inputDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := usecase.NewUpdateProductUseCase(h.productRepo)

	outputDTO, err := usecase.Execute(inputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(outputDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *UpdateProductHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("PUT /product", h.UpdateProductHandle)
}
