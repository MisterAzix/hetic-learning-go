package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"hetic-learning-go/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (productRepository *ProductRepository) FindAll() []model.Product {
	rows, err := productRepository.db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

func (productRepository *ProductRepository) FindById(id int) (model.Product, error) {
	row := productRepository.db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	product := model.Product{}
	err := row.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return product, fmt.Errorf("Product with ID %d not found", id)
		}
		return product, err
	}
	return product, nil
}

func (productRepository *ProductRepository) Save(product model.Product) model.Product {
	result, err := productRepository.db.Exec("INSERT INTO products (id, title, description, price, quantity) VALUES (?, ?, ?, ?, ?)", product.Id, product.Title, product.Description, product.Price, product.Quantity)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	product.Id = int(id)
	return product
}

func (productRepository *ProductRepository) Update(product model.Product) model.Product {
	_, err := productRepository.db.Exec("UPDATE products SET title = ?, description = ?, price = ?, quantity = ? WHERE id = ?", product.Title, product.Description, product.Price, product.Quantity, product.Id)
	if err != nil {
		panic(err)
	}
	return product
}
