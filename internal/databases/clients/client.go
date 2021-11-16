package clients

import "github.com/gna69/grpc-users/internal/databases/repo"

type (
	UsersClient interface {
		UsersService() repo.UserService
	}

	LogClient interface {
		LogService() repo.LogService
	}
)
