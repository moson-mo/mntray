package tray

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ClientSettings holds information about the max items and the categories to be fetched.
type ClientSettings struct {
	MaxItems   int      `json:"MaxItems"`
	Categories []string `json:"Categories"`
}

// connect to server and wait for responses
func (t *TrayIcon) waitForNews() {
	go func() {
		// prepare client settings
		cs := &ClientSettings{
			MaxItems:   t.Conf.MaxArticles,
			Categories: t.Conf.Categories,
		}

		// set categories according to branch if configured
		if t.Conf.SetCategoriesFromBranch {
			if t.Conf.userBranch == "" {
				fmt.Println("Could not detect branch, going for stable updates only")
				cs.Categories = []string{"Stable Updates", "Announcements"}
			} else {
				cs.Categories = []string{t.Conf.userBranch + " Updates", "Announcements"}
			}
		}

		b, _ := json.Marshal(cs)

		// wait for configured number of seconds if started in delay mode
		if t.Delay {
			<-time.After(time.Duration(t.Conf.DelayAfterStart) * time.Second)
		}
		t.addNewsFromServer(b)

		// fetch news loop (interval from config)
		for {
			<-time.After(time.Duration(t.Conf.RefreshInterval) * time.Second)
			t.addNewsFromServer(b)
		}
	}()
}

// adds news to menu
func (t *TrayIcon) addNewsFromServer(payload []byte) {
	// get articles from server
	articles, err := t.getArticlesFromServer(payload)
	if err != nil {
		fmt.Println("Error fetching news:", err)
		if t.Conf.ErrorNotifications {
			t.Icon.ErrorSlot(err)
		}
		return
	}
	// add news
	for _, a := range articles {
		t.gotNewArticle(a, false)
	}
}

// fetches articles from server
func (t *TrayIcon) getArticlesFromServer(body []byte) ([]Article, error) {
	// request news articles
	r, err := http.Post(t.Conf.ServerURL, "application/json", bytes.NewReader(body))
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
	articles := []Article{}
	err = json.Unmarshal(b, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
