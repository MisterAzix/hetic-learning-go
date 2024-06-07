package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const DBDriver = "mysql"
const DBUser = "hetic"
const DBPassword = "password"
const DBName = "hetic-learning-go"
const DBHost = "localhost"
const DBPort = "3306"

func InitDatabase() *sql.DB {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
	db, err := sql.Open(DBDriver, uri)
	if err != nil {
		panic(err)
	}
	migrate(*db)
	return db
}

func migrate(db sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY AUTO_INCREMENT, firstname VARCHAR(255), lastname VARCHAR(255), email VARCHAR(255), phone VARCHAR(255), address VARCHAR(255))")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id INT PRIMARY KEY AUTO_INCREMENT, title VARCHAR(255), description TEXT, price DECIMAL(10, 2), quantity INT)")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id INT PRIMARY KEY AUTO_INCREMENT, user_id INT, product_id INT, quantity INT, FOREIGN KEY (user_id) REFERENCES users(id), FOREIGN KEY (product_id) REFERENCES products(id))")
	if err != nil {
		panic(err)
	}
}
