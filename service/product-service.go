package service

import (
	"fmt"
	"hetic-learning-go/model"
	"hetic-learning-go/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
	orderRepository   repository.OrderRepository
}

func NewProductService(productRepository repository.ProductRepository, orderRepository repository.OrderRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		orderRepository:   orderRepository,
	}
}

func (productService *ProductService) AddInStock(title string, description string, price float64, quantity int) model.Product {
	product := model.Product{
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	return productService.productRepository.Save(product)
}

func (productService *ProductService) BuyProductById(productId int, quantity int, userId int) (model.Order, error) {
	product, err := productService.productRepository.FindById(productId)
	if err != nil {
		return model.Order{}, err
	}
	if product.Quantity <= 0 {
		return model.Order{}, fmt.Errorf("Product out of stock")
	}
	product.Quantity--
	productService.productRepository.Update(product)
	order := model.Order{
		UserId:     userId,
		ProductId:  productId,
		Quantity:   quantity,
		TotalPrice: product.Price * float64(quantity),
	}
	return productService.orderRepository.Save(order), nil
}

func (productService *ProductService) GetAll() []model.Product {
	return productService.productRepository.FindAll()
}

func (productService *ProductService) UpdateById(id int, title string, description string, price float64, quantity int) model.Product {
	product := model.Product{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	return productService.productRepository.Update(product)
}

func (productService *ProductService) DeactivateProduct(id int) model.Product {
	product, _ := productService.productRepository.FindById(id)
	product.IsActive = false
	return productService.productRepository.Update(product)
}
