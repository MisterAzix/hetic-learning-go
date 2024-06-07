package service

import (
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

func (productService *ProductService) BuyProductById(productId string, quantity int, userId string) model.Order {
	product := productService.productRepository.FindById(productId)
	if product.Quantity <= 0 {
		panic("Product out of stock")
	}
	product.Quantity--
	productService.productRepository.Update(product)
	order := model.Order{
		UserId:     userId,
		ProductId:  productId,
		Quantity:   quantity,
		TotalPrice: product.Price * float64(quantity),
	}
	return productService.orderRepository.Save(order)
}

func (productService *ProductService) GetAll() []model.Product {
	return productService.productRepository.FindAll()
}

func (productService *ProductService) UpdateById(id string, title string, description string, price float64, quantity int) model.Product {
	product := productService.productRepository.FindById(id)
	product.Title = title
	product.Description = description
	product.Price = price
	product.Quantity = quantity
	return productService.productRepository.Update(product)
}

func (productService *ProductService) DeactivateProduct(id string) model.Product {
	product := productService.productRepository.FindById(id)
	product.IsActive = false
	return productService.productRepository.Update(product)
}
