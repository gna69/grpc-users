package logger

import (
	"context"
	"fmt"
	"github.com/gna69/grpc-users/config"
	"github.com/gna69/grpc-users/internal/consts"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func Logger(kafkaConfig config.KafkaConfig) {
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
		fmt.Println(m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Debug("failed to close reader:", err)
	}
}
