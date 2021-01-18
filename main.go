package main

import (
	"github.com/handybots/toemoji/handler"
	tb "gopkg.in/tucnak/telebot.v3"
	"log"
	"os"
)

func main() {

	b, err := tb.NewBot(tb.Settings{Token: os.Getenv("TOKEN"), Poller: &tb.LongPoller{Timeout: 10}})
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.New(handler.Config{
		Bot: b,
	})

	b.Handle("/start", h.OnStart)
	//b.Handle(, h.OnStartTranslate)

	b.Handle(tb.OnText, h.OnText)
	b.Handle(tb.OnQuery, h.OnQuery)

	b.Poller = tb.NewMiddlewarePoller(b.Poller, h.OnUpdate)
	b.Start()
}
