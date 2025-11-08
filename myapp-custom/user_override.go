package main

import (
	"fmt"
	"myapp-main/core"
)

// Custom repository
type CustomRepository struct {
	data []core.User
}

func (r *CustomRepository) Save(user core.User) error {
	user.ID = len(r.data) + 100 // custom id offset
	fmt.Println("🔥 [CustomRepo] Saving user:", user.Name)
	r.data = append(r.data, user)
	return nil
}

func (r *CustomRepository) FindAll() []core.User {
	fmt.Println("🔥 [CustomRepo] Returning all users with custom logic")
	return r.data
}

// Custom service override
type CustomUserService struct {
	Repo core.UserRepository
}

func (s *CustomUserService) CreateUser(name string) error {
	fmt.Println("🔥 [CustomService] Custom validation before create:", name)
	user := core.User{Name: "[CUSTOM]" + name}
	return s.Repo.Save(user)
}

func (s *CustomUserService) GetUsers() []core.User {
	fmt.Println("🔥 [CustomService] Fetching users (custom)")
	return s.Repo.FindAll()
}

// init() otomatis dijalankan ketika di-run
func init() {
	fmt.Println("Registering custom overrides...")
	customRepo := &CustomRepository{}
	customService := &CustomUserService{Repo: customRepo}
	core.ActiveUserRepository = customRepo
	core.ActiveUserService = customService
}
