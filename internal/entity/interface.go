package entity

type ProductRepoInterface interface {
	CreateProduct(product *Product) error
	GetProducts() ([]*Product, error)
	GetProductByID(id string) (*Product, error)
	UpdateProduct(product *Product) (*Product, error)
	DeleteProduct(id string) error
}
