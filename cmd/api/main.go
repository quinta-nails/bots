package main

import (
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/bufbuild/protovalidate-go"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/quinta-nails/bots/internal/config"
	"github.com/quinta-nails/bots/internal/db"
	"github.com/quinta-nails/bots/internal/server"
	pb "github.com/quinta-nails/protobuf/gen/go/bots"
	pbCompanies "github.com/quinta-nails/protobuf/gen/go/companies"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Service struct {
	pb.UnimplementedTelegramBotsServiceServer
	db                     *db.Queries
	validator              *protovalidate.Validator
	companiesServiceClient pbCompanies.CompaniesServiceClient
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
	servicesConfig := config.ServicesConfig{}
	if err := env.Parse(&servicesConfig); err != nil {
		log.Fatal(err)
	}

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

	companiesServiceConnection, err := grpc.NewClient(servicesConfig.CompaniesEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't connect to companies grpc server")
	}
	defer companiesServiceConnection.Close()
	service.companiesServiceClient = pbCompanies.NewCompaniesServiceClient(companiesServiceConnection)

	grpcServer := grpc.NewServer()
	pb.RegisterTelegramBotsServiceServer(grpcServer, service)

	err = server.ListenAndServe(grpcServer)
	if err != nil {
		log.Fatal(err)
	}
}
