package bot

import (
	"strings"

	"github.com/handybots/toemoji"
	"github.com/handybots/toemoji/database"
	"github.com/handybots/toemoji/translate"

	"go.massbots.xyz/telebot/monitor"
	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
	"gopkg.in/tucnak/telebot.v3/middleware"
)

type Bot struct {
	*tele.Bot
	*layout.Layout
	db  *database.DB
	mon *monitor.Monitor
}

func New(path string, boot toemoji.Bootstrap) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot:    b,
		Layout: lt,
		db:     boot.DB,
		mon:    boot.Monitor,
	}, nil
}

func (b *Bot) Start() {
	b.OnError = b.mon.OnError()
	b.Use(middleware.DefaultLogger())
	b.Use(b.Middleware("ru"))

	b.Handle("/start", b.onStart)
	b.Handle(b.Callback("start_translate"), b.onStartTranslate)
	b.Handle(tele.OnText, b.onText)
	b.Handle(tele.OnQuery, b.onQuery)
}

func translateText(text string) (string, error) {
	result, err := translate.Translate(text)
	if err != nil {
		return "", err
	}
	return strings.Join(result.Text, ""), nil
}
