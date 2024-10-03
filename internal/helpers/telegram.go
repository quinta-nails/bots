package helpers

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-telegram/bot"
	"github.com/quinta-nails/bots/internal/config"
)

func NewTelegramBot(token string) (*bot.Bot, error) {
	telegramConfig := config.TelegramConfig{}
	if err := env.Parse(&telegramConfig); err != nil {
		return nil, err
	}

	botOptions := []bot.Option{
		bot.WithSkipGetMe(),
	}
	if telegramConfig.IsTestEnvironment {
		botOptions = append(botOptions, bot.UseTestEnvironment())
	}
	b, err := bot.New(token, botOptions...)
	if err != nil {
		return nil, err
	}

	return b, nil
}
