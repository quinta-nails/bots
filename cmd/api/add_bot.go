package main

import (
	"context"
	"github.com/quinta-nails/protobuf/gen/go/telegram_backend"
)

func (s *Service) AddBot(ctx context.Context, in *telegram_backend.AddBotRequest) (*telegram_backend.AddBotResponse, error) {
	resp := &telegram_backend.AddBotResponse{}

	err := s.validator.Validate(in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
