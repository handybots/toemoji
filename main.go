package main

import (
	"log"
	"os"

	tb "github.com/demget/telebot"
	"github.com/handybots/toemoji/handler"
)

func main() {
	tmplEngine := &tb.TemplateText{
		Dir:        "data",
		DelimLeft:  "${",
		DelimRight: "}",
	}

	pref, err := tb.NewSettingsYAML("bot.yaml", tmplEngine)
	if err != nil {
		log.Fatalln(err)
	}
	pref.Token = os.Getenv("TOKEN")

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.New(handler.Config{
		Bot: b,
	})

	b.Handle("/start", h.OnStart)
	b.Handle(b.InlineButton("start_translate"), h.OnStartTranslate)

	b.Handle(tb.OnText, h.OnText)
	b.Handle(tb.OnQuery, h.OnQuery)

	b.Poller = tb.NewMiddlewarePoller(b.Poller, h.OnUpdate)
	b.Start()
}
