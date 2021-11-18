package server

import (
	"github.com/gna69/grpc-users/config"
	usersDB "github.com/gna69/grpc-users/internal/api/users/db"
	"github.com/gna69/grpc-users/internal/api/users/handlers"
	"github.com/gna69/grpc-users/internal/databases/db/postgres"
	"github.com/gna69/grpc-users/internal/databases/db/redis"
	"github.com/gna69/grpc-users/internal/logger/broker"
)

func New(appConfig *config.Config) (*handlers.UserService, error) {
	db, err := postgres.New(appConfig.Postgres)
	if err != nil {
		return nil, err
	}

	usersClient, err := usersDB.New(db)
	if err != nil {
		return nil, err
	}

	kafkaClient, err := broker.New(appConfig.Kafka)
	if err != nil {
		return nil, err
	}

	return &handlers.UserService{
		UsersClient: usersClient,
		KafkaClient: kafkaClient,
		Cache:       redis.New(appConfig.Redis),
	}, nil
}
