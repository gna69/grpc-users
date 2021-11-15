package db

import (
	"github.com/gna69/grpc-users/internal/databases/db/postgres/user"
	"github.com/gna69/grpc-users/internal/databases/repo"
	"github.com/jinzhu/gorm"
)

type Client struct {
	db           *gorm.DB
	usersService repo.UserService
}

func New(db *gorm.DB) (*Client, error) {
	usersService, err := user.New(db)
	if err != nil {
		return nil, err
	}

	return &Client{
		db:           db,
		usersService: usersService,
	}, nil
}

func (c *Client) UsersService() repo.UserService {
	return c.usersService
}
