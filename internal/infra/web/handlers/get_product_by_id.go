package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/microservice_products/internal/entity"
	"github.com/NayronFerreira/microservice_products/internal/usecase"
)

type GetProductByIDHandler struct {
	productRepo entity.ProductRepoInterface
}

func NewGetProductByIDHandler(productRepo entity.ProductRepoInterface) *GetProductByIDHandler {
	return &GetProductByIDHandler{
		productRepo: productRepo,
	}
}

func (h *GetProductByIDHandler) GetProductByIDHandle(w http.ResponseWriter, r *http.Request) {

	var inputDTO usecase.GetProductByIDInputDTO
	if err := json.NewDecoder(r.Body).Decode(&inputDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if inputDTO.ID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	usecase := usecase.NewGetProductByIDUseCase(h.productRepo)

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

func (h *GetProductByIDHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /product", h.GetProductByIDHandle)
}
