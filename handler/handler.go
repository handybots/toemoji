package handler

import (
	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
)

func New(c Handler) handler {
	return handler{
		lt: c.Layout,
		b:  c.Bot,
	}
}

type (
	Handler struct {
		Layout *layout.Layout
		Bot    *tele.Bot
	}

	handler struct {
		lt *layout.Layout
		b  *tele.Bot
	}
)
