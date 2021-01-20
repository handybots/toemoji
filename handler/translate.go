package handler

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/handybots/toemoji/translate"

	"go.uber.org/atomic"
)

var currentSID atomic.String

func init() {
	data, err := ioutil.ReadFile("sid.txt")
	if err != nil {
		log.Println(err)
	}

	go func() {
		if data != nil {
			currentSID.Store(string(data))
			time.Sleep(24 * time.Hour)
		}
		for {
			sid, err := translate.ParseSID()
			if err != nil {
				log.Fatal(err)
			}

			currentSID.Store(sid)
			time.Sleep(24 * time.Hour)
		}
	}()
}

func translateText(text string) (string, error) {
	sid := currentSID.Load()
	if sid == "" {
		return "", errors.New("sid is empty")
	}
	result, err := translate.Translate(currentSID.Load(), text)
	if err != nil {
		return "", err
	}
	return strings.Join(result.Text, ""), nil
}
