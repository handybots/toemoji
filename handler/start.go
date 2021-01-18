package handler

import tb "gopkg.in/tucnak/telebot.v3"

func (h Handler) OnStart(context tb.Context) error {
	h.b.Send(context.Message().Sender,
		"start",
		"start",
		tb.NoPreview,
		tb.ModeMarkdown)
	return nil
}

func (h Handler) OnStartTranslate(context tb.Context) error {
	h.b.Edit(context.Callback().Message,
		"start",
		tb.NoPreview,
		tb.ModeMarkdown)
	return nil
}
