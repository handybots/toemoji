package handler

import (
	"log"

	tele "gopkg.in/tucnak/telebot.v3"
)

func (h handler) OnStart(c tele.Context) error {
	chat := c.Sender()

	exists, err := h.db.Users.Exists(chat)
	if err != nil {
		return err
	}

	if !exists {
		log.Println("Start from", chat.Recipient())
		if err := h.db.Users.Create(chat); err != nil {
			return err
		}
	}

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
