package tray

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

const (
	dir           = "/mntray"
	fileSettings  = "/settings.json"
	fileArticles  = "/articles.json"
	url           = "ws://manjaro.moson.eu:6666/articles"
	maxItems      = 10
	retryInterval = 10
	numberRetries = 5
	hideWhenRead  = false
)

// default categories
var categories = []string{"Testing Updates", "Stable Updates", "Unstable Updates", "Announcements", "manjaro32"}

// Settings to be saved to file
type settings struct {
	URL               string
	MaxArticles       int
	Categories        []string
	configDir         string
	configFile        string
	ArticlesFile      string
	ReconnectInterval int
	NumberOfRetries   int
	HideNoNews        bool
}

// NewSettings creates settings with default config if not yet existing
// Otherwise, load settings from file
func NewSettings() (*settings, error) {
	s := settings{
		URL:               url,
		MaxArticles:       maxItems,
		Categories:        categories,
		ReconnectInterval: retryInterval,
		NumberOfRetries:   numberRetries,
		HideNoNews:        hideWhenRead,
	}
	d, err := createConfigDir()
	if err != nil {
		return nil, err
	}
	s.configDir = d
	s.configFile = d + fileSettings
	s.ArticlesFile = d + fileArticles

	err = s.LoadSettings()
	if err != nil {
		fmt.Println("No config found. Trying to create default config file")
		err = s.SaveSettings()
		if err != nil {
			return nil, err
		}
	}

	return &s, nil
}

// Save saves settings to file
func (s *settings) SaveSettings() error {
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	ioutil.WriteFile(s.configFile, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Save saves articles to file
func (s *settings) SaveArticles(items []article) error {
	b, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return err
	}

	ioutil.WriteFile(s.ArticlesFile, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Load loads setting from file
func (s *settings) LoadSettings() error {
	b, err := ioutil.ReadFile(s.configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, s)
	if err != nil {
		return err
	}
	return nil
}

// Load loads articles from file
func (s *settings) LoadArticles() ([]article, error) {
	b, err := ioutil.ReadFile(s.ArticlesFile)
	if err != nil {
		return nil, err
	}

	articles := &[]article{}
	err = json.Unmarshal(b, articles)
	if err != nil {
		return nil, err
	}

	sort.Slice(*articles, func(i, j int) bool {
		return (*articles)[i].PublishedDate.Unix() < (*articles)[j].PublishedDate.Unix()
	})

	return *articles, nil
}

// creates the config dir if not yet existing
func createConfigDir() (string, error) {
	d, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	d += "/.config" + dir

	_, err = os.Stat(d)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}
		err = os.Mkdir(d, 0755)
		if err != nil {
			return "", err
		}
	}
	return d, nil
}
