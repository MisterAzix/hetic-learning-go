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
		err := writer.Write([]string{user.Id, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address})
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
		newQuantity := strconv.Itoa(product.Quantity)
		newPrice := strconv.FormatFloat(product.Price, 'f', -1, 64)
		err := writer.Write([]string{product.Id, product.Title, product.Description, newQuantity, newPrice})
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
		newQuantity := strconv.Itoa(order.Quantity)
		newTotalPrice := strconv.FormatFloat(order.TotalPrice, 'f', -1, 64)
		err := writer.Write([]string{order.Id, order.UserId, order.ProductId, newQuantity, newTotalPrice})
		if err != nil {
			panic(err)
		}
	}

	return nil
}
