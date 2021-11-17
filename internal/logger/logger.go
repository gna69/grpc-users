package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gna69/grpc-users/config"
	"github.com/gna69/grpc-users/internal/consts"
	"github.com/gna69/grpc-users/internal/databases/db/clickhouse"
	"github.com/gna69/grpc-users/internal/databases/model"
	"github.com/gna69/grpc-users/internal/logger/client"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"time"
)

func castBytesToUser(bytesUser []byte) (*model.User, error) {
	var user *model.User
	err := json.Unmarshal(bytesUser, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Logger(kafkaConfig config.KafkaConfig, clickHouse config.ClickHouseConfig) {
	db, err := clickhouse.New(clickHouse)
	if err != nil {
		log.Debug("error connect to clickhouse db: ", err.Error())
		return
	}

	chClient, err := client.New(db)
	if err != nil {
		log.Debug("error getting clickhouse service: ", err.Error())
		return
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s:%d", kafkaConfig.Host, kafkaConfig.Port)},
		Topic:     consts.Topic,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e2,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		user, err := castBytesToUser(m.Value)
		if err != nil {
			log.Debugf("error unmarshal user: %s, error: %s", m.Value, err.Error())
			continue
		}

		logRecord := &model.LogRecord{
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			UserId:     uint32(user.Id),
			ActionDate: time.Now(),
			ActionTime: time.Now(),
		}

		err = chClient.LogService().Create(logRecord)
		if err != nil {
			log.Debug("error writing log to clickhouse: ", err.Error())
		}

		fmt.Println(m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Debug("failed to close reader:", err)
	}
}
