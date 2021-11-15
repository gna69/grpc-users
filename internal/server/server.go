package server

import (
	api "github.com/gna69/grpc-users/pkg/api"
)

func New() *api.UserService {
	srv := &api.UserService{}

	return srv
}
