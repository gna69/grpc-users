package handlers

import (
	"context"
	"encoding/json"
	"github.com/gna69/grpc-users/internal/databases/clients"
	"github.com/gna69/grpc-users/internal/databases/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	UsersClient clients.UsersClient
	KafkaClient *kafka.Writer
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

func castUserStructToBytes(user *model.User) ([]byte, error) {
	buffer, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func (s *UserService) Add(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	castedUser := castProtoToLocal(user)

	err := s.UsersClient.UsersService().Create(ctx, castedUser)
	if err != nil {
		return nil, err
	}

	bytesUser, err := castUserStructToBytes(castedUser)
	if err == nil {
		err = s.KafkaClient.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte("user"),
			Value: bytesUser,
		})
		if err != nil {
			log.Debug("error sending kafka message: ", err.Error())
		}
	} else {
		log.Debug("error convert struct to []byte: ", err.Error())
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
