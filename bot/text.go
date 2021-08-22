package bot

import (
	"strings"

	tele "gopkg.in/tucnak/telebot.v3"
)

func (b Bot) onText(c tele.Context) error {
	if strings.Contains(c.Text(), "\n") {
		return nil
	}

	result, err := translateText(c.Text())
	if err != nil {
		return err
	}

	return c.Reply(result)
}
