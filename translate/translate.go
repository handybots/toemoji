package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	urlMain         = "https://translate.yandex.com"
	urlApiBase      = "https://translate.yandex.net/api"
	urlApiTranslate = urlApiBase + "/v1/tr.json/translate?"

	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"
)

var defaultParams = url.Values{
	"srv":    {"tr-text"},
	"lang":   {"ru-emj"},
	"reason": {"auto"},
	"format": {"text"},
}

type Result struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Lang    string   `json:"lang"`
	Text    []string `json:"text"`
}

func Translate(sid, text string) (Result, error) {
	params := url.Values{}
	params.Set("id", sid)
	for k, v := range defaultParams {
		params[k] = v
	}

	form := url.Values{}
	form.Set("text", text)
	form.Set("option", "4")

	endp := urlApiTranslate + params.Encode()
	body := strings.NewReader(form.Encode())

	req, err := http.NewRequest(http.MethodPost, endp, body)
	if err != nil {
		return Result{}, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Result{}, fmt.Errorf("translate: response code is %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{}, err
	}

	var result Result
	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}
	if result.Code != 200 {
		return result, fmt.Errorf("translate: result code is %d %s", result.Code, result.Message)
	}
	return result, nil
}
