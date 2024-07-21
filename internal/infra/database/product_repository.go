package database

import (
	"database/sql"
	"fmt"

	"github.com/NayronFerreira/microservice_products/configs"
	"github.com/NayronFerreira/microservice_products/internal/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (d *ProductRepository) CreateProduct(product *entity.Product) error {
	_, err := d.db.Exec("INSERT INTO products (id, name, model, code, price, color) VALUES (?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Model, product.Code, product.Price, product.Color)
	if err != nil {
		return err
	}
	return nil
}

func (d *ProductRepository) GetProducts() ([]*entity.Product, error) {
	rows, err := d.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Model, &product.Code, &product.Price, &product.Color)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (d *ProductRepository) GetProductByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := d.db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Model, &product.Code, &product.Price, &product.Color); err != nil {
		return nil, err
	}
	return &product, nil
}

func (d *ProductRepository) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	stmt, err := d.db.Prepare("UPDATE products SET name=?, model=?, code=?, price=?, color=? WHERE id=?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.Name, product.Model, product.Code, product.Price, product.Color, product.ID)
	if err != nil {
		return nil, err
	}

	updatedProduct, err := d.GetProductByID(product.ID)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (d *ProductRepository) DeleteProduct(id string) error {
	_, err := d.db.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (d *ProductRepository) SetupTable(db *sql.DB, configs configs.Config) error {

	query := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        id VARCHAR(255) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        model VARCHAR(255) NOT NULL,
        code VARCHAR(50) NOT NULL,
        price DECIMAL(10,2) NOT NULL,
        color VARCHAR(50) NOT NULL
    );
    `, configs.DBTable)

	_, err := d.db.Exec(query)
	return err
}
