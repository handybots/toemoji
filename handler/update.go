package handler

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v3"
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
		user = u.Query.Sender
		data = u.Query.Text
	default:
		return true
	}

	log.Println(user.ID, data)
	return true
}
