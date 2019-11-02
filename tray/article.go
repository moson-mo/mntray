package tray

import (
	"time"

	"github.com/therecipe/qt/widgets"
)

// Article struct. Holds information for a news article
type Article struct {
	GUID          string    `json:"GUID"`
	URL           string    `json:"URL"`
	Title         string    `json:"Title"`
	Category      string    `json:"Category"`
	PublishedDate time.Time `json:"PublishedDate"`
	WasRead       bool      `json:"WasRead"`
	mi            *widgets.QAction
}
