package tray

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// defaults
const (
	dir               = "/mntray"
	fileSettings      = "/settings.json"
	fileArticles      = "/articles.json"
	url               = "http://manjaro.moson.eu:10111/news"
	maxItems          = 10
	refreshInterval   = 600
	hideWhenRead      = false
	autostart         = true
	errorNotification = true
	desktopFile       = `[Desktop Entry]
Terminal=false
Name=mntray
Type=Application
Exec=/usr/bin/mntray
Icon=mntray
Comment=A Manjaro Linux announcements notification app
`
)

// default categories
var categories = []string{"Testing Updates", "Stable Updates", "Unstable Updates", "Announcements", "manjaro32"}

// Settings to be saved to file
type settings struct {
	ServerURL          string
	MaxArticles        int
	Categories         []string
	configDir          string
	configBaseDir      string
	configFile         string
	ArticlesFile       string
	RefreshInterval    int
	HideNoNews         bool
	Autostart          bool
	ErrorNotifications bool
}

// NewSettings creates settings with default config if not yet existing
// Otherwise, load settings from file
func NewSettings() (*settings, error) {
	s := settings{
		ServerURL:          url,
		MaxArticles:        maxItems,
		Categories:         categories,
		RefreshInterval:    refreshInterval,
		HideNoNews:         hideWhenRead,
		Autostart:          autostart,
		ErrorNotifications: errorNotification,
	}
	d, err := createConfigDir()
	if err != nil {
		return nil, err
	}
	s.configDir = d
	s.configBaseDir = strings.Replace(d, "/mntray", "", 1)
	s.configFile = d + fileSettings
	s.ArticlesFile = d + fileArticles

	err = s.LoadSettings()
	if err != nil {
		fmt.Println("No config found. Trying to create default config file")
		err = s.SaveSettings(false)
		if err != nil {
			return nil, err
		}
	}

	err = s.createDesktopFile()
	if err != nil {
		print(err)
	}

	return &s, nil
}

// Save saves settings to file
func (s *settings) SaveSettings(loadBeforeSave bool) error {
	if loadBeforeSave {
		err := s.LoadSettings()
		if err != nil {
			print("Settings could not be loaded: " + err.Error())
		}
	}
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	ioutil.WriteFile(s.configFile, b, 0644)
	if err != nil {
		return err
	}

	err = s.createDesktopFile()
	if err != nil {
		return err
	}
	return nil
}

func (s *settings) createDesktopFile() error {
	asdir := s.configBaseDir + "/autostart"
	fileName := asdir + "/mntray.desktop"

	// make sure we have an autostart dir (might be useless?)
	_, err := os.Stat(asdir)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		err = os.Mkdir(asdir, 0644)
		if err != nil {
			return err
		}
	}

	// create file (if not yet existing)
	if s.Autostart {
		_, err := os.Stat(fileName)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}

			err = ioutil.WriteFile(fileName, []byte(desktopFile), 0755)
			if err != nil {
				return err
			}
		}
	} else { // remove file if existing
		_, err := os.Stat(fileName)
		if err == nil {
			print("remove desktop file")
			err = os.Remove(fileName)
			if err != nil {
				return err
			}
		}
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
