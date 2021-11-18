package user

import (
	"context"
	"github.com/gna69/grpc-users/internal/databases/model"
	"go.elastic.co/apm/module/apmgorm"
)

func (s *service) Create(ctx context.Context, user *model.User) error {
	db := apmgorm.WithContext(ctx, s.db)
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(ctx context.Context, user *model.User) (*model.User, error) {
	db := apmgorm.WithContext(ctx, s.db)
	if err := db.Where("id = ?", user.Id).Delete(&model.User{}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetAll(ctx context.Context) ([]*model.User, error) {
	db := apmgorm.WithContext(ctx, s.db)
	var users []*model.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
