package handler

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v3"
)

func (h Handler) OnText(context tb.Context) error {
	result, err := translateText(context.Text())
	if err != nil {
		log.Println(err)
		return err
	}

	h.b.Reply(context.Message(), result)
	return nil
}
