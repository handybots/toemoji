package bot

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b Bot) onQuery(c tele.Context) error {
	text := c.Data()
	if text == "" {
		text = "привет"
	}

	result, err := translateText(text)
	if err != nil {
		return err
	}

	return c.Answer(&tele.QueryResponse{
		Results:           tele.Results{b.Result(c, "tr", result)},
		SwitchPMText:      b.String("switch_pm_text"),
		SwitchPMParameter: b.String("switch_pm_param"),
		CacheTime:         1000,
	})
}
