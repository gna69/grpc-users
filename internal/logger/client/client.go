package client

import (
	logs "github.com/gna69/grpc-users/internal/databases/db/clickhouse/log"
	"github.com/gna69/grpc-users/internal/databases/repo"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	logsService repo.LogService
}

func New(db *sqlx.DB) (*Client, error) {
	logsService, err := logs.New(db)
	if err != nil {
		return nil, err
	}

	return &Client{
		logsService: logsService,
	}, nil
}

func (c *Client) LogService() repo.LogService {
	return c.logsService
}
