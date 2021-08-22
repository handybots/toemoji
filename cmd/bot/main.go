package main

import (
	"log"
	"os"

	"github.com/handybots/toemoji"
	"github.com/handybots/toemoji/bot"
	"github.com/handybots/toemoji/database"
	"github.com/handybots/toemoji/translate"

	"go.massbots.xyz/telebot/monitor"
)

func main() {
	db, err := database.Open(os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	mon, err := monitor.New(monitorConfig)
	if err != nil {
		log.Fatal(err)
	}

	boot := toemoji.Bootstrap{
		DB:      db,
		Monitor: mon,
	}

	b, err := bot.New("bot.yml", boot)
	if err != nil {
		log.Fatal(err)
	}

	// Workers
	go translate.WatchSID(b.Duration("translate_d"))

	b.Start()
}

var monitorConfig = monitor.Config{
	URL: os.Getenv("CLICKHOUSE_URL"),
}
