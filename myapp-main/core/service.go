package core

import "fmt"

type UserService interface {
	CreateUser(name string) error
	GetUsers() []User
}

type DefaultUserService struct {
	Repo UserRepository
}

func (s *DefaultUserService) CreateUser(name string) error {
	fmt.Println("[Core] CreateUser called with:", name)
	user := User{Name: name}
	return s.Repo.Save(user)
}

func (s *DefaultUserService) GetUsers() []User {
	fmt.Println("[Core] GetUsers called")
	return s.Repo.FindAll()
}

// variabel global untuk di-override
var ActiveUserService UserService = &DefaultUserService{
	Repo: GetUserRepository(),
}

func GetUserService() UserService {
	return ActiveUserService
}
