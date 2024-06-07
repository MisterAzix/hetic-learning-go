package service

import (
	"encoding/csv"
	"hetic-learning-go/repository"
	"os"
	"strconv"
)

type CSVService struct {
	csvFile *os.File
	repository.UserRepository
	repository.ProductRepository
	repository.OrderRepository
}

func NewCSVService(userRepository repository.UserRepository, productRepository repository.ProductRepository, orderRepository repository.OrderRepository) *CSVService {
	return &CSVService{
		UserRepository:    userRepository,
		ProductRepository: productRepository,
		OrderRepository:   orderRepository,
	}
}

func (csvService *CSVService) ExportAllUsers() error {
	writer := csv.NewWriter(csvService.csvFile)
	defer writer.Flush()

	users := csvService.UserRepository.FindAll()
	for _, user := range users {
		id := strconv.Itoa(user.Id)
		err := writer.Write([]string{id, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address})
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (csvService *CSVService) ExportAllProducts() error {
	writer := csv.NewWriter(csvService.csvFile)
	defer writer.Flush()

	products := csvService.ProductRepository.FindAll()
	for _, product := range products {
		id := strconv.Itoa(product.Id)
		quantity := strconv.Itoa(product.Quantity)
		price := strconv.FormatFloat(product.Price, 'f', -1, 64)
		err := writer.Write([]string{id, product.Title, product.Description, quantity, price})
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (csvService *CSVService) ExportAllOrders() error {
	writer := csv.NewWriter(csvService.csvFile)
	defer writer.Flush()

	orders := csvService.OrderRepository.FindAll()
	for _, order := range orders {
		id := strconv.Itoa(order.Id)
		userId := strconv.Itoa(order.UserId)
		productId := strconv.Itoa(order.ProductId)
		quantity := strconv.Itoa(order.Quantity)
		price := strconv.FormatFloat(order.TotalPrice, 'f', -1, 64)
		err := writer.Write([]string{id, userId, productId, quantity, price})
		if err != nil {
			panic(err)
		}
	}

	return nil
}
