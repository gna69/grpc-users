package broker

import (
	"fmt"
	"github.com/gna69/grpc-users/config"
	"github.com/gna69/grpc-users/internal/consts"
	"github.com/segmentio/kafka-go"
)

func New(config config.KafkaConfig) (*kafka.Writer, error) {
	return &kafka.Writer{
		Addr:     kafka.TCP(fmt.Sprintf("%s:%d", config.Host, config.Port)),
		Topic:    consts.Topic,
		Balancer: &kafka.LeastBytes{},
	}, nil
}
