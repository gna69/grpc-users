package repo

import (
	"context"
	"github.com/gna69/grpc-users/internal/databases/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}
