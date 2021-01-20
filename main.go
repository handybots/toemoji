package main

import (
	"log"
	"os"

	"github.com/handybots/toemoji/handler"
	"github.com/handybots/toemoji/storage"

	"github.com/demget/clickrus"
	"github.com/sirupsen/logrus"

	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
	"gopkg.in/tucnak/telebot.v3/middleware"
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

	db, err := storage.Open(os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	ch, err := clickrus.NewHook(clickHouseConfig)
	if err != nil {
		// log.Fatal(err)
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	_ = ch // logger.AddHook(ch)

	h := handler.New(handler.Handler{
		Layout: lt,
		Bot:    b,
		DB:     db,
	})

	b.OnError = h.OnError
	b.Use(middleware.Logger(logger, h.LoggerFields))
	b.Use(lt.Middleware("ru"))

	b.Handle("/start", h.OnStart)
	b.Handle(lt.Callback("start_translate"), h.OnStartTranslate)
	b.Handle(tele.OnText, h.OnText)
	b.Handle(tele.OnQuery, h.OnQuery)

	b.Start()
}

var clickHouseConfig = clickrus.Config{
	Addr:    os.Getenv("CLICKHOUSE_URL"),
	Columns: []string{"event", "user_id"},
	Table:   "toemoji.logs",
}
