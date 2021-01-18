package handler

import tb "github.com/demget/telebot"

func (h Handler) OnStart(m *tb.Message) {
	h.b.Send(m.Sender,
		h.b.Text("start"),
		h.b.InlineMarkup("start"),
		tb.NoPreview,
		tb.ModeMarkdown)
}

func (h Handler) OnStartTranslate(c *tb.Callback) {
	h.b.Edit(c.Message,
		h.b.Text("start_ru"),
		tb.NoPreview,
		tb.ModeMarkdown)
}
