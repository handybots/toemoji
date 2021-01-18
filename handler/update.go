package handler

import (
	"log"

	tb "github.com/demget/telebot"
)

func (h Handler) OnUpdate(u *tb.Update) bool {
	var (
		user *tb.User
		data string
	)
	switch {
	case u.Message != nil:
		user = u.Message.Sender
		data = u.Message.Text
	case u.Query != nil:
		user = &u.Query.From
		data = u.Query.Text
	default:
		return true
	}

	log.Println(user.ID, data)
	return true
}
