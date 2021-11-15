package handlers

import (
	"context"
	"github.com/gna69/grpc-users/internal/databases/clients"
	"github.com/gna69/grpc-users/internal/databases/model"
)

type UserService struct {
	UsersClient clients.UsersClient
}

func castProtoToLocal(user *UserInfo) *model.User {
	return &model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func castLocalToProto(user *model.User) *UserInfo {
	return &UserInfo{
		Id:        int32(user.Id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func (s *UserService) Add(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	castedUser := castProtoToLocal(user)

	err := s.UsersClient.UsersService().Create(ctx, castedUser)
	if err != nil {
		return nil, err
	}

	newUser := castLocalToProto(castedUser)

	return newUser, nil
}

func (s *UserService) Remove(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	removedUser := castProtoToLocal(user)

	removedUser, err := s.UsersClient.UsersService().Delete(ctx, removedUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, e *Empty) (*AllUsers, error) {

	users, err := s.UsersClient.UsersService().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var allUsers []*UserInfo
	for _, user := range users {
		grpcUser := castLocalToProto(user)
		allUsers = append(allUsers, grpcUser)
	}

	return &AllUsers{Users: allUsers}, nil
}

func (s *UserService) mustEmbedUnimplementedUserServer() {
	panic("implement me")
}
