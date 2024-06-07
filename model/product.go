package model

type Product struct {
	Id          string
	Title       string
	Description string
	Price       float64
	Quantity    int
	IsActive    bool
}

func NewProduct(id string, title string, description string, price float64, quantity int) *Product {
	return &Product{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}
