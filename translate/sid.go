package translate

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"go.uber.org/atomic"
)

var (
	reSID = regexp.MustCompile(`SID: *'([^']+)'`)
)

var currentSID atomic.String

func WatchSID(d time.Duration) {
	data, err := ioutil.ReadFile("sid.txt")
	if err != nil {
		log.Println(err)
	}

	if data != nil {
		currentSID.Store(string(data))
		time.Sleep(d)
	}

	for {
		sid, err := parseSID()
		if err != nil {
			log.Fatal(err)
		}

		currentSID.Store(sid)
		time.Sleep(d)
	}
}

func parseSID() (string, error) {
	req, err := http.NewRequest(http.MethodGet, urlMain, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("translate: response code is %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	match := reSID.FindStringSubmatch(string(data))
	if len(match) != 2 {
		return "", ErrNoSID
	}

	group := strings.Split(match[1], ".")
	if len(group) == 0 {
		return "", ErrNoSID
	}

	for i, g := range group {
		group[i] = reverseString(g)
	}

	return strings.Join(group, ".") + "-0-0", nil
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
