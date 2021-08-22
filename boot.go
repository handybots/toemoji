package toemoji

import (
	"github.com/handybots/toemoji/database"
	"go.massbots.xyz/telebot/monitor"
)

type Bootstrap struct {
	DB      *database.DB
	Monitor *monitor.Monitor
}
