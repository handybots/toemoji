package handler

import (
	tb "github.com/demget/telebot"
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
