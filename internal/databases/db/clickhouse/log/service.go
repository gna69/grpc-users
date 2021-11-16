package log

import (
	"github.com/gna69/grpc-users/internal/databases/repo"
	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

var schema = `create table  if not exist logs (
    							id UInt32,
                                first_name String,
                                last_name String,
                                user_id UInt32,
                                action_day Date,
                                action_time DateTime
                                ) engine Memory`

func New(dbConn *sqlx.DB) (repo.LogService, error) {
	lService := &service{db: dbConn}

	if err := lService.CreateLogsTableIfNotExist(); err != nil {
		return nil, err
	}

	return lService, nil
}

func (s *service) CreateLogsTableIfNotExist() error {
	s.db.MustExec(schema)

	return nil
}
