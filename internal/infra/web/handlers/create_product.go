package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/microservice_products/internal/entity"
	"github.com/NayronFerreira/microservice_products/internal/usecase"
)

type CreateProductHandler struct {
	productRepo entity.ProductRepoInterface
}

func NewCreateProductHandler(productRepo entity.ProductRepoInterface) *CreateProductHandler {
	return &CreateProductHandler{productRepo: productRepo}
}

func (h CreateProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var createProductInputDTO usecase.CreateProductInputDTO
	if err := json.NewDecoder(r.Body).Decode(&createProductInputDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productDTO, err := usecase.NewCreateProductUseCase(h.productRepo).Execute(createProductInputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(productDTO)
}

func (h *CreateProductHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /product", h.CreateProduct)
}
