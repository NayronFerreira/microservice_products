package usecase

import "github.com/NayronFerreira/microservice_products/internal/entity"

type UpdateProductInputDTO struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
}

type UpdateProductOutputDTO struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
	Color string  `json:"color"`
}

type UpdateProductUseCase struct {
	productRepo entity.ProductRepoInterface
}

func NewUpdateProductUseCase(productRepo entity.ProductRepoInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepo: productRepo,
	}
}

func (u *UpdateProductUseCase) Execute(productInpuDTO UpdateProductInputDTO) (UpdateProductOutputDTO, error) {
	product := &entity.Product{
		ID:    productInpuDTO.ID,
		Name:  productInpuDTO.Name,
		Model: productInpuDTO.Model,
		Code:  productInpuDTO.Code,
		Price: productInpuDTO.Price,
		Color: productInpuDTO.Color,
	}

	product, err := u.productRepo.UpdateProduct(product)
	if err != nil {
		return UpdateProductOutputDTO{}, err
	}

	return UpdateProductOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Model: product.Model,
		Code:  product.Code,
		Price: product.Price,
		Color: product.Color,
	}, nil
}
