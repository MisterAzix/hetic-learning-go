package service

import (
	"hetic-learning-go/model"
	"hetic-learning-go/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) Register(firstname string, lastname string, email string, phone string, address string) model.User {
	user := model.User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Phone:     phone,
		Address:   address,
	}
	return userService.userRepository.Save(user)
}

func (userService *UserService) GetAll() []model.User {
	return userService.userRepository.FindAll()
}

func (userService *UserService) UpdateById(id int, firstname string, lastname string, email string, phone string, address string) model.User {
	user := model.User{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Phone:     phone,
		Address:   address,
	}
	return userService.userRepository.UpdateById(user)
}
