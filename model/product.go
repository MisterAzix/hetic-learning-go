package model

type Product struct {
	Id          int
	Title       string
	Description string
	Price       float64
	Quantity    int
	IsActive    bool
}

func NewProduct(id int, title string, description string, price float64, quantity int) *Product {
	return &Product{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}
