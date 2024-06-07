package service

import (
	"hetic-learning-go/model"
	"hetic-learning-go/repository"
)

type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (orderService *OrderService) GetAll() []model.Order {
	return orderService.orderRepository.FindAll()
}
