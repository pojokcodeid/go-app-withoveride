package core

import "fmt"

type UserRepository interface {
	Save(user User) error
	FindAll() []User
}

type InMemoryUserRepository struct {
	data []User
}

func (r *InMemoryUserRepository) Save(user User) error {
	user.ID = len(r.data) + 1
	fmt.Println("[Core] Saving user:", user.Name)
	r.data = append(r.data, user)
	return nil
}

func (r *InMemoryUserRepository) FindAll() []User {
	fmt.Println("[Core] Retrieving all users")
	return r.data
}

var ActiveUserRepository UserRepository = &InMemoryUserRepository{}

func GetUserRepository() UserRepository {
	return ActiveUserRepository
}
