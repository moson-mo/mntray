package tray

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// defaults
const (
	version           = "0.2.0"
	dir               = "/mntray"
	fileConfig        = "/settings.json"
	fileArticles      = "/articles.json"
	url               = "http://manjaro.moson.eu:10111/news"
	maxItems          = 10
	refreshInterval   = 600
	hideWhenRead      = false
	autostart         = true
	errorNotification = true
	fetchDelay        = 60
	branchDetect      = true
	desktopFile       = `[Desktop Entry]
Terminal=false
Name=mntray
Type=Application
Exec=/usr/bin/mntray --delay
Icon=mntray
Comment=A Manjaro Linux announcements notification app
`
)

// default categories
var categories = []string{"Testing Updates", "Stable Updates", "Unstable Updates", "Announcements", "manjaro32"}

// Config to be saved to file
type Config struct {
	ServerURL               string
	MaxArticles             int
	Categories              []string
	configDir               string
	configBaseDir           string
	configFile              string
	articlesFile            string
	RefreshInterval         int
	HideNoNews              bool
	Autostart               bool
	ErrorNotifications      bool
	DelayAfterStart         int
	SetCategoriesFromBranch bool
	Version                 string
	userBranch              string
}

// NewConfig creates Config with default config if not yet existing
// Otherwise, load Config from file
func NewConfig() (*Config, error) {
	s := Config{
		ServerURL:               url,
		MaxArticles:             maxItems,
		Categories:              categories,
		RefreshInterval:         refreshInterval,
		HideNoNews:              hideWhenRead,
		Autostart:               autostart,
		ErrorNotifications:      errorNotification,
		DelayAfterStart:         fetchDelay,
		SetCategoriesFromBranch: branchDetect,
	}
	d, err := createConfigDir()
	if err != nil {
		return nil, err
	}
	s.configDir = d
	s.configBaseDir = strings.Replace(d, "/mntray", "", 1)
	s.configFile = d + fileConfig
	s.articlesFile = d + fileArticles
	s.userBranch = getBranch()

	err = s.LoadConfig()
	if err != nil {
		fmt.Println("No config found. Trying to create default config file")
		err = s.SaveConfig(false)
		if err != nil {
			return nil, err
		}
	}
	replaceDesktopFile := false

	if s.Version == "" {
		replaceDesktopFile = true
	}

	s.Version = version

	err = s.createDesktopFile(replaceDesktopFile)
	if err != nil {
		print(err)
	}

	return &s, nil
}

// SaveConfig saves Config to file
func (s *Config) SaveConfig(loadBeforeSave bool) error {
	if loadBeforeSave {
		err := s.LoadConfig()
		if err != nil {
			print("Config could not be loaded: " + err.Error())
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

	err = s.createDesktopFile(false)
	if err != nil {
		return err
	}
	return nil
}

func (s *Config) createDesktopFile(replace bool) error {
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
		// delete previous desktop file (needed for migration to v0.2.0)
		if replace {
			os.Remove(fileName) // we don't care about errors
		}
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
			fmt.Println("removing desktop file")
			err = os.Remove(fileName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SaveArticles saves Articles to file
func (s *Config) SaveArticles(items []Article) error {
	b, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return err
	}

	ioutil.WriteFile(s.articlesFile, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfig loads setting from file
func (s *Config) LoadConfig() error {
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

// LoadArticles loads Articles from file
func (s *Config) LoadArticles() ([]Article, error) {
	b, err := ioutil.ReadFile(s.articlesFile)
	if err != nil {
		return nil, err
	}

	Articles := &[]Article{}
	err = json.Unmarshal(b, Articles)
	if err != nil {
		return nil, err
	}

	sort.Slice(*Articles, func(i, j int) bool {
		return (*Articles)[i].PublishedDate.Unix() < (*Articles)[j].PublishedDate.Unix()
	})

	return *Articles, nil
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

// get branch from pacman-mirrors config file
func getBranch() string {
	f, err := os.Open("/etc/pacman-mirrors.conf")
	if err != nil {
		fmt.Println("Error getting branch:", err)
		return ""
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := strings.ToLower(s.Text())
		if strings.Contains(line, "branch") && len(line) > 0 && line[0] != '#' {
			b := strings.Split(strings.Replace(line, " ", "", -1), "=")
			if len(b) > 1 {
				return strings.Title(b[1])
			}
			return ""
		}
	}
	return ""

}
