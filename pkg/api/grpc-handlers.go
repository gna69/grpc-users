package api

import (
	"context"
)

type UserService struct{}

func (s *UserService) Add(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	return user, nil
}

func (s *UserService) Remove(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, e *Empty) (*AllUsers, error) {
	return nil, nil
}

func (s *UserService) mustEmbedUnimplementedUserServer() {
	panic("implement me")
}
