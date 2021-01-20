package handler

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (h handler) OnQuery(c tele.Context) error {
	text := c.Data()
	if text == "" {
		text = "привет"
	}

	result, err := translateText(text)
	if err != nil {
		return err
	}

	view := struct {
		Text   string
		Result string
	}{
		Text:   text,
		Result: result,
	}

	r := &tele.ArticleResult{
		Title:       view.Text,
		Description: view.Result,
	}
	r.SetContent(&tele.InputTextMessageContent{
		Text:      h.lt.Text(c, "result", view),
		ParseMode: tele.ModeMarkdown,
	})

	return c.Answer(&tele.QueryResponse{
		Results:           tele.Results{r},
		SwitchPMText:      h.lt.String("switch_pm_text"),
		SwitchPMParameter: h.lt.String("switch_pm_param"),
		CacheTime:         1000,
	})
}
