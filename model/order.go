package model

import (
	"time"
)

type Order struct {
	Id         int
	UserId     int
	ProductId  int
	Quantity   int
	TotalPrice float64
	OrderAt    time.Time
}

func NewOrder(id int, userId int, productId int, quantity int, totalPrice float64) *Order {
	return &Order{
		Id:         id,
		UserId:     userId,
		ProductId:  productId,
		Quantity:   quantity,
		TotalPrice: totalPrice,
		OrderAt:    time.Now(),
	}
}
