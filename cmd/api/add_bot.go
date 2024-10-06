package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-telegram/bot"
	"github.com/quinta-nails/bots/internal/db"
	"github.com/quinta-nails/bots/internal/helpers"
	pb "github.com/quinta-nails/protobuf/gen/go/bots"
)

func (s *Service) AddBot(ctx context.Context, in *pb.AddBotRequest) (*pb.AddBotResponse, error) {
	resp := &pb.AddBotResponse{}

	err := s.validator.Validate(in)
	if err != nil {
		return nil, err
	}

	botModel, err := s.db.GetBotByToken(ctx, in.Token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if err == nil {
		return nil, errors.New("bot already exists")
	}

	b, err := helpers.NewTelegramBot(in.Token)
	if err != nil {
		return nil, err
	}

	telegramBot, err := b.GetMe(ctx)
	if err != nil && errors.Is(err, bot.ErrorUnauthorized) {
		return nil, errors.New("invalid token")
	}
	if err != nil {
		return nil, err
	}

	botModel, err = s.db.AddBot(ctx, db.AddBotParams{
		Token:     in.Token,
		FirstName: telegramBot.FirstName,
		Username:  telegramBot.Username,
	})
	if err != nil {
		return nil, err
	}

	resp.Result = helpers.NewPbBotFromBotRow(&botModel)

	return resp, nil
}
