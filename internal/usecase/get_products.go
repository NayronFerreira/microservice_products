package usecase

import "github.com/NayronFerreira/microservice_products/internal/entity"

type GetProductsOutputDTO struct {
	Products []ProductDTO `json:"products"`
}

type ProductDTO struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
}

type GetProductsUseCase struct {
	productRepo entity.ProductRepoInterface
}

func NewGetProductsUseCase(productRepo entity.ProductRepoInterface) *GetProductsUseCase {
	return &GetProductsUseCase{productRepo: productRepo}
}

func (u *GetProductsUseCase) Execute() (*GetProductsOutputDTO, error) {
	products, err := u.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	var productsDTO []ProductDTO
	for _, product := range products {
		productDTO := ProductDTO{
			ID:    product.ID,
			Name:  product.Name,
			Model: product.Model,
			Code:  product.Code,
			Price: product.Price,
			Color: product.Color,
		}
		productsDTO = append(productsDTO, productDTO)
	}

	return &GetProductsOutputDTO{Products: productsDTO}, nil
}
