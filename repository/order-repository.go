package repository

import (
	"database/sql"
	"hetic-learning-go/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (orderRepository *OrderRepository) FindAll() []model.Order {
	rows, err := orderRepository.db.Query("SELECT * FROM orders")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		order := model.Order{}
		err := rows.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.TotalPrice, &order.OrderAt)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	return orders
}

func (orderRepository *OrderRepository) FindById(id string) model.Order {
	row := orderRepository.db.QueryRow("SELECT * FROM orders WHERE id = ?", id)
	order := model.Order{}
	err := row.Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.TotalPrice, &order.OrderAt)
	if err != nil {
		panic(err)
	}
	return order
}

func (orderRepository *OrderRepository) Save(order model.Order) model.Order {
	result, err := orderRepository.db.Exec("INSERT INTO orders (id, userId, productId, quantity, totalPrice, orderAt) VALUES (?, ?, ?, ?, ?, ?)", order.Id, order.UserId, order.ProductId, order.Quantity, order.TotalPrice, order.OrderAt)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	order.Id = int(id)
	return order
}
