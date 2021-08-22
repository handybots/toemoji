package bot

import (
	"log"

	tele "gopkg.in/tucnak/telebot.v3"
)

func (b Bot) onStart(c tele.Context) error {
	chat := c.Sender()

	exists, err := b.db.Users.Exists(chat)
	if err != nil {
		return err
	}

	if !exists {
		log.Println("Start from", chat.Recipient())
		if err := b.db.Users.Create(chat); err != nil {
			return err
		}
	}

	return c.Send(
		b.Text(c, "start"),
		b.Markup(c, "start"),
		tele.NoPreview,
	)
}

func (b Bot) onStartTranslate(c tele.Context) error {
	return c.Edit(
		b.Text(c, "start_translate"),
		b.Markup(c, "switch_inline"),
		tele.NoPreview,
	)
}
