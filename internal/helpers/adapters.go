package helpers

import (
	"github.com/quinta-nails/protobuf/gen/go/telegram_backend"
	"github.com/quinta-nails/telegram-backend/internal/db"
)

func NewPbBotFromBotRow(row *db.Bot) *telegram_backend.Bot {
	return &telegram_backend.Bot{
		Id:        row.ID,
		FirstName: row.FirstName,
		Username:  row.Username,
	}
}
