package model

type User struct {
	Id        int
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Address   string
}

func NewUser(id int, firstname string, lastname string, email string, phone string, address string) *User {
	return &User{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Phone:     phone,
		Address:   address,
	}
}
