package main

import (
	"log"

	"github.com/handybots/toemoji/handler"

	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
)

func main() {
	lt, err := layout.New("bot.yml")
	if err != nil {
		log.Fatal(err)
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(handler.Handler{
		Layout: lt,
		Bot:    b,
	})

	b.Use(lt.Middleware("ru"))
	b.Handle("/start", h.OnStart)
	b.Handle(lt.Callback("start_translate"), h.OnStartTranslate)
	b.Handle(tele.OnText, h.OnText)
	b.Handle(tele.OnQuery, h.OnQuery)

	b.Start()
}
