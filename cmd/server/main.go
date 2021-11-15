package main

import (
	"fmt"
	"github.com/gna69/grpc-users/config"
	"github.com/gna69/grpc-users/internal/api/users/handlers"
	"github.com/gna69/grpc-users/internal/server"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func init() {
	if _, ok := os.LookupEnv("DEBUG"); ok {
		err := godotenv.Load("./config/dev.env")
		if err != nil {
			log.Fatal("Error loading dev.env file")
		}
	}

	if _, ok := os.LookupEnv("DEBUG"); !ok {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	appConfig := config.Get()

	srv, err := server.New(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	handlers.RegisterUserServer(grpcServer, srv)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.Port))
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Server started on %d port", appConfig.Port)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
