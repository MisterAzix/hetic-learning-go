package main

import (
	"fmt"
	"hetic-learning-go/database"
	"hetic-learning-go/repository"
	"hetic-learning-go/service"
	"log"
)

func main() {
	log.Println("Starting application...")
	db := database.InitDatabase()
	log.Println("Database connection established!")

	log.Println("Initializing repositories...")
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(*orderRepository)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(*productRepository, *orderRepository)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	log.Println("Repositories initialized!")

	orders := orderService.GetAll()
	products := productService.GetAll()
	users := userService.GetAll()

	fmt.Printf("Orders: %v\n", orders)
	fmt.Printf("Products: %v\n", products)
	fmt.Printf("Users: %v\n", users)
}
