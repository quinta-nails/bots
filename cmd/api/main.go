package main

import (
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/bufbuild/protovalidate-go"
	"github.com/joho/godotenv"
	pb "github.com/quinta-nails/protobuf/gen/go/telegram_backend"
	"github.com/quinta-nails/telegram-backend/internal/db"
	"github.com/quinta-nails/telegram-backend/internal/server"
	"google.golang.org/grpc"
	"log"
)

type Service struct {
	pb.UnimplementedTelegramBackendServiceServer
	db        *db.Queries
	validator *protovalidate.Validator
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error loading .env file`)
	}

	err = db.ApplyMigrations()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	service := &Service{}

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}
	service.validator = validator

	dbConnection, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	service.db = dbConnection

	grpcServer := grpc.NewServer()
	pb.RegisterTelegramBackendServiceServer(grpcServer, service)

	err = server.ListenAndServe(grpcServer)
	if err != nil {
		log.Fatal(err)
	}
}
