package user

import (
	"context"
	"github.com/gna69/grpc-users/internal/databases/model"
	"github.com/gna69/grpc-users/internal/databases/repo"
	"github.com/jinzhu/gorm"
	"go.elastic.co/apm/module/apmgorm"
)

type service struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) (repo.UserService, error) {
	uService := &service{db: dbConn}

	if err := uService.CreateUsersTableIfNotExist(context.Background()); err != nil {
		return nil, err
	}

	return uService, nil
}

func (s *service) CreateUsersTableIfNotExist(ctx context.Context) error {
	db := apmgorm.WithContext(ctx, s.db)

	db.AutoMigrate(&model.User{})
	if !db.HasTable(&model.User{}) {
		if err := db.CreateTable(&model.User{}).Error; err != nil {
			return err
		}
	}

	return nil
}
