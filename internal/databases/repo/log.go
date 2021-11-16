package repo

import "github.com/gna69/grpc-users/internal/databases/model"

type LogService interface {
	Create(log *model.LogRecord) error
}
