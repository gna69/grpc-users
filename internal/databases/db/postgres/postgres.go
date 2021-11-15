package postgres

import (
	"fmt"
	"github.com/gna69/grpc-users/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	Postgres string = "postgres"
)

func New(config config.PostgresConfig) (*gorm.DB, error) {
	dbAddr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Database, config.Password)

	db, err := gorm.Open(Postgres, dbAddr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
