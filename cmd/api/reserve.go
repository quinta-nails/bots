package main

import (
	"context"
	"database/sql"
	"errors"
	pb "github.com/quinta-nails/protobuf/gen/go/bots"
	pbCompanies "github.com/quinta-nails/protobuf/gen/go/companies"
)

func (s *Service) Reserve(ctx context.Context, in *pb.ReserveRequest) (*pb.ReserveResponse, error) {
	resp := &pb.ReserveResponse{}

	err := s.validator.Validate(in)
	if err != nil {
		return nil, err
	}

	studioId, err := s.db.GetStudioIdByBotId(ctx, in.BotId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("studio not found")
	}
	if err != nil {
		return nil, err
	}

	_, err = s.companiesServiceClient.Reserve(ctx, &pbCompanies.ReserveRequest{
		StudioId:  studioId.Int64,
		UserId:    in.UserId,
		ServiceId: in.ServiceId,
		Datetime:  in.Datetime,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
