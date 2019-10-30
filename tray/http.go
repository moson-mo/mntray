package tray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type clientSettings struct {
	MaxItems   int      `json:"MaxItems"`
	Categories []string `json:"Categories"`
}

// connect to server and wait for responses
func (ti *trayIcon) waitForNews() {
	go func() {
		ti.wg.Add(1)
		defer ti.wg.Done()
		// prepare client settings
		cs := &clientSettings{
			MaxItems:   ti.config.MaxArticles,
			Categories: ti.config.Categories,
		}
		b, err := json.Marshal(cs)
		if err != nil {
			print(err)
			return
		}

		select {
		case <-time.After(time.Duration(ti.config.DelayAfterStart) * time.Second):
			fmt.Println("fetched news...")
			ti.addNewsFromServer(b)
		case <-ti.quit:
			return
		}

		// wait for news to arrive from server
		for {
			select {
			case <-time.After(time.Duration(ti.config.RefreshInterval) * time.Second):
				ti.addNewsFromServer(b)
			case <-ti.quit:
				return
			}
		}
	}()
}

func (ti *trayIcon) addNewsFromServer(payload []byte) {
	// get articles from server
	articles, err := ti.getArticlesFromServer(payload)
	if err != nil {
		print(err)
		if ti.config.ErrorNotifications {
			ti.icon.ConnectionDead(err)
		}
		return
	}
	// add news
	for _, a := range articles {
		ti.receivedNews(a, false)
	}
}

// get articles from server
func (ti *trayIcon) getArticlesFromServer(body []byte) ([]article, error) {
	// request news articles
	r, err := http.Post(ti.config.ServerURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// get data
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// get articles from json
	articles := []article{}
	err = json.Unmarshal(b, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
