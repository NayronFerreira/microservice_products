package usecase

import (
	"github.com/NayronFerreira/microservice_products/internal/entity"
	"github.com/google/uuid"
)

type CreateProductInputDTO struct {
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
}

type CreateProductOutputDTO struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
}

type CreateProductUseCase struct {
	productRepo entity.ProductRepoInterface
}

func NewCreateProductUseCase(productRepo entity.ProductRepoInterface) *CreateProductUseCase {
	return &CreateProductUseCase{productRepo: productRepo}
}

func (u *CreateProductUseCase) Execute(input CreateProductInputDTO) (*CreateProductOutputDTO, error) {
	id := uuid.New().String()
	product := entity.Product{
		ID:    id,
		Name:  input.Name,
		Model: input.Model,
		Code:  input.Code,
		Price: input.Price,
		Color: input.Color,
	}

	if err := u.productRepo.CreateProduct(&product); err != nil {
		return nil, err
	}

	return &CreateProductOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Model: product.Model,
		Code:  product.Code,
		Price: product.Price,
		Color: product.Color,
	}, nil
}
