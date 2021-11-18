package handlers

import (
	"context"
	"encoding/json"
	"github.com/gna69/grpc-users/internal/databases/clients"
	"github.com/gna69/grpc-users/internal/databases/model"
	"github.com/go-redis/cache/v8"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserService struct {
	UsersClient clients.UsersClient
	KafkaClient *kafka.Writer
	Cache       *cache.Cache
}

const (
	CachedKey = "users"
)

func castProtoToLocal(user *UserInfo) *model.User {
	return &model.User{
		Id:        uint(user.Id),
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

func (s *UserService) updateCache(ctx context.Context, users []*UserInfo) error {
	return s.Cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   CachedKey,
		Value: users,
		TTL:   time.Minute,
	})
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

	var cachedUsers []*UserInfo
	if err := s.Cache.Get(ctx, CachedKey, &cachedUsers); err == nil {
		cachedUsers = append(cachedUsers, newUser)
		err = s.updateCache(ctx, cachedUsers)
		if err != nil {
			log.Debug("cant update cache after creating new user: ", err.Error())
		}
	}

	return newUser, nil
}

func (s *UserService) Remove(ctx context.Context, user *UserInfo) (*UserInfo, error) {
	removedUser := castProtoToLocal(user)

	removedUser, err := s.UsersClient.UsersService().Delete(ctx, removedUser)
	if err != nil {
		return nil, err
	}

	var cachedUsers []*UserInfo
	if err := s.Cache.Get(ctx, CachedKey, &cachedUsers); err == nil {
		for i, cachedUser := range cachedUsers {
			if user.Id == cachedUser.Id {
				cachedUsers[i] = cachedUsers[len(cachedUsers)-1]
				cachedUsers[len(cachedUsers)-1] = nil
				cachedUsers = cachedUsers[:len(cachedUsers)-1]
				err = s.updateCache(ctx, cachedUsers)
				if err != nil {
					log.Debug("cant update cache after deleting user: ", err.Error())
				}
				break
			}
		}
	}

	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, e *Empty) (*AllUsers, error) {
	var cachedUsers []*UserInfo
	if err := s.Cache.Get(ctx, CachedKey, &cachedUsers); err == nil {
		return &AllUsers{Users: cachedUsers}, nil
	}

	users, err := s.UsersClient.UsersService().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var allUsers []*UserInfo
	for _, user := range users {
		grpcUser := castLocalToProto(user)
		allUsers = append(allUsers, grpcUser)
	}

	if err := s.updateCache(ctx, allUsers); err != nil {
		log.Debug("error caching users to redis: ", err.Error())
	}

	return &AllUsers{Users: allUsers}, nil
}

func (s *UserService) mustEmbedUnimplementedUserServer() {
	panic("implement me")
}
