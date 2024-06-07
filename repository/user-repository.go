package repository

import (
	"database/sql"
	"hetic-learning-go/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepository *UserRepository) FindAll() []model.User {
	rows, err := userRepository.db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func (userRepository *UserRepository) FindById(id string) model.User {
	row := userRepository.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	user := model.User{}
	err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address)
	if err != nil {
		panic(err)
	}
	return user
}

func (userRepository *UserRepository) Save(user model.User) model.User {
	result, err := userRepository.db.Exec("INSERT INTO users (id, firstname, lastname, email, phone, address) VALUES (?, ?, ?, ?, ?, ?)", user.Id, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.Id = int(id)
	return user
}

func (userRepository *UserRepository) UpdateById(user model.User) model.User {
	_, err := userRepository.db.Exec("UPDATE users SET firstname = ?, lastname = ?, email = ?, phone = ?, address = ? WHERE id = ?", user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.Id)
	if err != nil {
		panic(err)
	}
	return user
}
