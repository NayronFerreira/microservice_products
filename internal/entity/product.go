package entity

type Product struct {
	ID    string
	Name  string
	Model string
	Code  string
	Price float64
	Color string
}

func NewProduct(id string, name, model, code string, price float64, color string) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Model: model,
		Code:  code,
		Price: price,
		Color: color,
	}
}
