package handler

import (
	"log"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v3"
)

func (h Handler) OnQuery(context tb.Context) error {
	var rs tb.Results

	if context.Query().Text != "" {
		result, err := translateText(context.Query().Text)
		if err != nil {
			log.Println(err)
			return err
		}

		view := struct {
			Text   string
			Result string
		}{
			Text:   context.Query().Text,
			Result: result,
		}

		r := &tb.ArticleResult{
			Title:       view.Text,
			Description: view.Result,
			Text:        view.Text + "\n" + view.Result,
		}
		r.SetResultID(strconv.Itoa(1))

		rs = append(rs, r)
	}

	err := h.b.Answer(context.Query(), &tb.QueryResponse{
		Results:           rs,
		SwitchPMText:      "Перейти в чат с ботом",
		SwitchPMParameter: "inline",
		CacheTime:         1000,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
