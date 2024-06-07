package model

import (
	"time"
)

type Order struct {
	Id         int
	UserId     string
	ProductId  string
	Quantity   int
	TotalPrice float64
	OrderAt    time.Time
}

func NewOrder(id int, userId string, productId string, quantity int, totalPrice float64) *Order {
	return &Order{
		Id:         id,
		UserId:     userId,
		ProductId:  productId,
		Quantity:   quantity,
		TotalPrice: totalPrice,
		OrderAt:    time.Now(),
	}
}
