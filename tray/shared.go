package tray

import (
	"fmt"
	"time"
)

type article struct {
	GUID          string    `json:"GUID"`
	URL           string    `json:"URL"`
	Title         string    `json:"Title"`
	Category      string    `json:"Category"`
	PublishedDate time.Time `json:"PublishedDate"`
	WasRead       bool      `json:"WasRead"`
}

// prints out a message with the current timestamp
func print(msg interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 - 15:04:05:"), msg)
}
