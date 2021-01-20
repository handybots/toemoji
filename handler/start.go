package handler

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (h handler) OnStart(c tele.Context) error {
	return c.Send(
		h.lt.Text(c, "start"),
		h.lt.Markup(c, "start"),
		tele.NoPreview,
	)
}

func (h handler) OnStartTranslate(c tele.Context) error {
	return c.Edit(
		h.lt.Text(c, "start_translate"),
		h.lt.Markup(c, "switch_inline"),
		tele.NoPreview,
	)
}
