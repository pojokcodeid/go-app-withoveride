package plugins

import "myapp-main/core"

// Inisialisasi pointer agar bisa diubah dari luar
var (
	RegisteredService    core.UserService
	RegisteredRepository core.UserRepository
)

// Fungsi akses
func GetService() core.UserService {
	if RegisteredService != nil {
		return RegisteredService
	}
	return core.GetUserService()
}

func GetRepository() core.UserRepository {
	if RegisteredRepository != nil {
		return RegisteredRepository
	}
	return core.GetUserRepository()
}
