package main

import "database/sql"

func main() {
	db, err := sql.Open("mysql", "root:root@tcp()")
}
