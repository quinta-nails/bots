package helpers

import (
	"github.com/quinta-nails/bots/internal/db"
	pb "github.com/quinta-nails/protobuf/gen/go/bots"
)

func NewPbBotFromBotRow(row *db.Bot) *pb.Bot {
	return &pb.Bot{
		Id:        row.ID,
		FirstName: row.FirstName,
		Username:  row.Username,
	}
}
