package clickhouse

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/gna69/grpc-users/config"
	"github.com/jmoiron/sqlx"
)

const (
	ClickHouse string = "clickhouse"
)

func New(config config.ClickHouseConfig) (*sqlx.DB, error) {
	dbAddr := fmt.Sprintf("tcp://%s:%d?debug=true", config.Host, config.Port)

	db, err := sqlx.Open(ClickHouse, dbAddr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
