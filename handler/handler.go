package handler

import (
	tb "gopkg.in/tucnak/telebot.v3"
)

type Handler struct {
	b *tb.Bot
}

type Config struct {
	Bot *tb.Bot
}

func New(c Config) Handler {
	return Handler{
		b: c.Bot,
	}
}
