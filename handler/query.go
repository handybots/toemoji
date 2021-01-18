package handler

import (
	"log"

	tb "github.com/demget/telebot"
)

func (h Handler) OnQuery(q *tb.Query) {
	var rs tb.Results

	if q.Text != "" {
		result, err := translateText(q.Text)
		if err != nil {
			log.Println(err)
			return
		}

		view := struct {
			Text   string
			Result string
		}{
			Text:   q.Text,
			Result: result,
		}

		r := h.b.InlineResult("result", view)
		r.SetContent(&tb.InputTextMessageContent{
			Text:      h.b.Text("result", view),
			ParseMode: tb.ModeMarkdown,
		})

		rs = append(rs, r)
	}

	err := h.b.Answer(q, &tb.QueryResponse{
		Results:           rs,
		SwitchPMText:      h.b.String("switch_pm_text"),
		SwitchPMParameter: h.b.String("switch_pm_param"),
		CacheTime:         1000,
	})
	if err != nil {
		log.Println(err)
		return
	}
}
