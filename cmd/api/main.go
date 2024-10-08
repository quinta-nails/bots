package main

import (
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/bufbuild/protovalidate-go"
	"github.com/joho/godotenv"
	"github.com/quinta-nails/bots/internal/db"
	"github.com/quinta-nails/bots/internal/server"
	pb "github.com/quinta-nails/protobuf/gen/go/bots"
	"google.golang.org/grpc"
	"log"
)

type Service struct {
	pb.UnimplementedTelegramBotsServiceServer
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
	pb.RegisterTelegramBotsServiceServer(grpcServer, service)

	err = server.ListenAndServe(grpcServer)
	if err != nil {
		log.Fatal(err)
	}
}
