package usecase

import "github.com/NayronFerreira/microservice_products/internal/entity"

type DeleteProductInputDTO struct {
	ID string `json:"id"`
}

type DeleteProductUseCase struct {
	productRepo entity.ProductRepoInterface
}

func NewDeleteProductUseCase(productRepo entity.ProductRepoInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productRepo: productRepo,
	}
}

func (u *DeleteProductUseCase) Execute(productInpuDTO DeleteProductInputDTO) error {
	err := u.productRepo.DeleteProduct(productInpuDTO.ID)
	if err != nil {
		return err
	}
	return nil
}
