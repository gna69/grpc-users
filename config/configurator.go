package config

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	Port       uint
	Postgres   PostgresConfig
	Kafka      KafkaConfig
	ClickHouse ClickHouseConfig
	Redis      RedisConfig
}

type PostgresConfig struct {
	Host     string
	Port     uint
	Database string
	User     string
	Password string
}

type KafkaConfig struct {
	Host string
	Port uint
}

type ClickHouseConfig struct {
	Host string
	Port uint
}

type RedisConfig struct {
	Host string
	Port uint
}

var config *Config

func Get() *Config {
	if config == nil {
		return readConfig()
	}
	return config
}

func readConfig() *Config {
	envValue, ok := os.LookupEnv("PORT")
	appPort, err := strconv.ParseInt(envValue, 10, 0)

	if !ok || err != nil {
		log.Warn("No app port in ENV. Using default 8080")
		appPort = 8080
	}

	envValue, ok = os.LookupEnv("PG_PORT")
	pgPort, err := strconv.ParseInt(envValue, 10, 0)

	if !ok || err != nil {
		log.Warn("No postgres db port in ENV. Using default 5432")
		pgPort = 5432
	}

	envValue, ok = os.LookupEnv("KAFKA_PORT")
	kafkaPort, err := strconv.ParseInt(envValue, 10, 0)

	if !ok || err != nil {
		log.Warn("No kafka port in ENV. Using default 9092")
		kafkaPort = 9092
	}

	envValue, ok = os.LookupEnv("CLICKHOUSE_PORT")
	clickhousePort, err := strconv.ParseInt(envValue, 10, 0)

	if !ok || err != nil {
		log.Warn("No clickhouse port in ENV. Using default 8123")
		clickhousePort = 8123
	}

	envValue, ok = os.LookupEnv("REDIS_PORT")
	redisPort, err := strconv.ParseInt(envValue, 10, 0)

	if !ok || err != nil {
		log.Warn("No redis port in ENV. Using default 6379")
		clickhousePort = 8123
	}

	config = &Config{
		Port: uint(appPort),
		Postgres: PostgresConfig{
			Host:     os.Getenv("PG_HOST"),
			Port:     uint(pgPort),
			Database: os.Getenv("PG_DB"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASS"),
		},
		Kafka: KafkaConfig{
			Host: os.Getenv("KAFKA_HOST"),
			Port: uint(kafkaPort),
		},
		ClickHouse: ClickHouseConfig{
			Host: os.Getenv("CLICKHOUSE_HOST"),
			Port: uint(clickhousePort),
		},
		Redis: RedisConfig{
			Host: os.Getenv("REDIS_HOST"),
			Port: uint(redisPort),
		},
	}

	return config
}
