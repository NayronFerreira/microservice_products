package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/microservice_products/internal/entity"
	"github.com/NayronFerreira/microservice_products/internal/usecase"
)

type GetAllProductsHandler struct {
	productRepo entity.ProductRepoInterface
}

func NewGetAllProductsHandler(productRepo entity.ProductRepoInterface) *GetAllProductsHandler {
	return &GetAllProductsHandler{productRepo: productRepo}
}

func (h GetAllProductsHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := usecase.NewGetProductsUseCase(h.productRepo).Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *GetAllProductsHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /products", h.GetAllProducts)
}
