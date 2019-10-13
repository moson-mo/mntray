package tray

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

var con *websocket.Conn

type clientSettings struct {
	MaxItems   int      `json:"MaxItems"`
	Categories []string `json:"Categories"`
}

type article struct {
	GUID          string    `json:"GUID"`
	URL           string    `json:"URL"`
	Title         string    `json:"Title"`
	Categories    []string  `json:"Categories"`
	PublishedDate time.Time `json:"PublishedDate"`
	WasRead       bool      `json:"WasRead"`
}

// connect to server and wait for responses
func (ti *trayIcon) waitForNews() {
	go func() {
		// try to connect, if not successfull, retry every according to settings
		cnt := 1
		for {
			err := ti.connect()
			if err != nil {
				fmt.Println("Error connecting: " + err.Error())
				time.Sleep(time.Duration(ti.config.ReconnectInterval) * time.Second)
			} else {
				break
			}

			if cnt >= ti.config.NumberOfRetries {
				fmt.Println("Can't connect to server, giving up...")
				ti.icon.ConnectionDead()
				return
			}
			cnt++
		}

		// wait for news to arrive from server
		for {
			a := &article{}

			err := con.ReadJSON(a)
			if err != nil {
				fmt.Println("Error: " + err.Error())
				ti.waitForNews()
				return
			}
			ti.receivedNews(*a, false)
		}
	}()
}

// connect to server
func (ti *trayIcon) connect() error {
	var err error

	con, _, err = websocket.DefaultDialer.Dial(ti.config.URL, nil)
	if err != nil {
		return err
	}

	s := clientSettings{
		Categories: ti.config.Categories,
		MaxItems:   ti.config.MaxArticles,
	}
	err = con.WriteJSON(s)
	if err != nil {
		return err
	}

	return nil
}
