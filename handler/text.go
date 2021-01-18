package handler

import (
	"log"

	tb "github.com/demget/telebot"
)

func (h Handler) OnText(m *tb.Message) {
	result, err := translateText(m.Text)
	if err != nil {
		log.Println(err)
		return
	}

	h.b.Reply(m, result)
}
