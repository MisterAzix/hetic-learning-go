package main

import (
	"fmt"
	"hetic-learning-go/database"
	"hetic-learning-go/handler"
	"hetic-learning-go/repository"
	"hetic-learning-go/service"
	"hetic-learning-go/utils"
	"log"
	"os"
)

func displayMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Add product")
	fmt.Println("2. Display products")
	fmt.Println("3. Update product")
	fmt.Println("4. Delete product")
	fmt.Println("5. Add user")
	fmt.Println("6. Display users")
	fmt.Println("7. Update user")
	fmt.Println("8. Buy product")

	fmt.Println("0. Exit")
}

func main() {
	log.Println("Starting application...")
	db := database.InitDatabase()
	log.Println("Database connection established!")

	log.Println("Initializing repositories...")
	orderRepository := repository.NewOrderRepository(db)
	productRepository := repository.NewProductRepository(db)
	userRepository := repository.NewUserRepository(db)
	log.Println("Repositories initialized!")

	log.Println("Initializing services...")
	orderService := service.NewOrderService(*orderRepository)
	productService := service.NewProductService(*productRepository, *orderRepository)
	userService := service.NewUserService(*userRepository)
	log.Println("Services initialized!")

	appHandler := handler.NewHandler(*orderService, *productService, *userService)

	for {
		displayMenu()

		var choice = utils.ReadInt("\nEnter your choice:")
		switch choice {
		case 1:
			appHandler.AddProduct()
		case 2:
			appHandler.DisplayProducts()
		case 3:
			appHandler.UpdateProduct()
		case 4:
			appHandler.DeleteProduct()
		case 5:
			appHandler.AddUser()
		case 6:
			appHandler.DisplayUsers()
		case 7:
			appHandler.UpdateUser()
		case 8:
			appHandler.BuyProduct()
		case 0:
			os.Exit(0)
		}
	}
}
