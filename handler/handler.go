package handler

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"hetic-learning-go/service"
	"hetic-learning-go/utils"
	"os"
)

type Handler struct {
	CSVService     service.CSVService
	OrderService   service.OrderService
	ProductService service.ProductService
	UserService    service.UserService
}

func NewHandler(csvService service.CSVService, orderService service.OrderService, productService service.ProductService, userService service.UserService) *Handler {
	return &Handler{
		CSVService:     csvService,
		OrderService:   orderService,
		ProductService: productService,
		UserService:    userService,
	}
}

func (handler *Handler) AddProduct() {
	title := utils.ReadInput("Enter product title:")
	description := utils.ReadInput("Enter product description:")
	quantity := utils.ReadFloat("Enter product quantity:")
	price := utils.ReadInt("Enter product price:")

	handler.ProductService.AddInStock(title, description, quantity, price)
}

func (handler *Handler) DisplayProducts() {
	products := handler.ProductService.GetAll()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title", "Description", "Quantity", "Price"})
	for _, product := range products {
		t.AppendRow(table.Row{product.Id, product.Title, product.Description, product.Quantity, product.Price})
	}
	t.Render()
}

func (handler *Handler) UpdateProduct() {
	id := utils.ReadInt("Enter product ID:")
	title := utils.ReadInput("Enter product title:")
	description := utils.ReadInput("Enter product description:")
	quantity := utils.ReadFloat("Enter product quantity:")
	price := utils.ReadInt("Enter product price:")

	handler.ProductService.UpdateById(id, title, description, quantity, price)
}

func (handler *Handler) DeleteProduct() {
	id := utils.ReadInt("Enter product ID:")
	handler.ProductService.DeactivateProduct(id)
}

func (handler *Handler) ExportProductsCSV() {
	err := handler.CSVService.ExportAllProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Products exported successfully!")
}

func (handler *Handler) AddUser() {
	firstname := utils.ReadInput("Enter user firstname:")
	lastname := utils.ReadInput("Enter user lastname:")
	email := utils.ReadInput("Enter user email:")
	phone := utils.ReadInput("Enter user phone:")
	address := utils.ReadInput("Enter user address:")

	handler.UserService.Register(firstname, lastname, email, phone, address)
}

func (handler *Handler) DisplayUsers() {
	users := handler.UserService.GetAll()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Firstname", "Lastname", "Email", "Phone", "Address"})
	for _, user := range users {
		t.AppendRow(table.Row{user.Id, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address})
	}
	t.Render()
}

func (handler *Handler) UpdateUser() {
	id := utils.ReadInt("Enter user ID:")
	firstname := utils.ReadInput("Enter user firstname:")
	lastname := utils.ReadInput("Enter user lastname:")
	email := utils.ReadInput("Enter user email:")
	phone := utils.ReadInput("Enter user phone:")
	address := utils.ReadInput("Enter user address:")

	handler.UserService.UpdateById(id, firstname, lastname, email, phone, address)
}

func (handler *Handler) ExportUsersCSV() {
	err := handler.CSVService.ExportAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Users exported successfully!")
}

func (handler *Handler) BuyProduct() {
	productId := utils.ReadInt("Enter product ID:")
	quantity := utils.ReadInt("Enter quantity:")
	userId := utils.ReadInt("Enter user ID:")

	_, err := handler.ProductService.BuyProductById(productId, quantity, userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Product bought successfully!")
}

func (handler *Handler) ExportOrdersCSV() {
	err := handler.CSVService.ExportAllOrders()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Orders exported successfully!")
}
