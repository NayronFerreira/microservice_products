package usecase

import (
	"errors"

	"github.com/NayronFerreira/microservice_products/internal/entity"
)

type GetProductByIDInputDTO struct {
	ID string `json:"id"`
}

type GetProductByIDOutputDTO struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
}

type GetProductByIDUseCase struct {
	productRepo entity.ProductRepoInterface
}

func NewGetProductByIDUseCase(productRepo entity.ProductRepoInterface) *GetProductByIDUseCase {
	return &GetProductByIDUseCase{
		productRepo: productRepo,
	}
}

func (u *GetProductByIDUseCase) Execute(productInpuDTO GetProductByIDInputDTO) (retVal GetProductByIDOutputDTO, err error) {
	product, err := u.productRepo.GetProductByID(productInpuDTO.ID)
	if err != nil {
		return retVal, err
	}

	if product.ID == "" {
		return retVal, errors.New("product not found")
	}

	return GetProductByIDOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Model: product.Model,
		Code:  product.Code,
		Price: product.Price,
		Color: product.Color,
	}, nil
}
