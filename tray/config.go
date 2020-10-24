package tray

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

// defaults
const (
	version           = "1.1.3"
	dir               = "/mntray"
	fileConfig        = "/settings.json"
	fileArticles      = "/articles.json"
	url               = "http://manjaro.moson.eu:10111/news"
	maxItems          = 15
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
	noAuto = "Hidden=true"
)

// default categories
var categories = []string{"All"}
var availableCategories = []string{"Testing Updates", "Stable Updates", "Stable Staging Updates", "Unstable Updates", "Twitter", "News", "Releases", "ARM News", "ARM Releases", "ARM Stable Updates", "ARM Testing Updates", "ARM Unstable Updates"}
var addCategoriesBranch = []string{"Announcements"}

// Config to be saved to file
type Config struct {
	Version                 string
	ServerURL               string
	MaxArticles             int
	AvailableCategories     []string
	Categories              []string
	AddCategoriesBranch     []string
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
	userBranch              string
	autostartChanged        bool
}

// NewConfig creates Config with default config if not yet existing
// Otherwise, load Config from file
func NewConfig() (*Config, error) {
	s := Config{
		ServerURL:               url,
		MaxArticles:             maxItems,
		Categories:              categories,
		AvailableCategories:     append(availableCategories[:0:0], availableCategories...),
		AddCategoriesBranch:     addCategoriesBranch,
		RefreshInterval:         refreshInterval,
		HideNoNews:              hideWhenRead,
		Autostart:               autostart,
		ErrorNotifications:      errorNotification,
		DelayAfterStart:         fetchDelay,
		SetCategoriesFromBranch: branchDetect,
		autostartChanged:        false,
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
	if s.userBranch == "" {
		fmt.Println("Could not detect branch, assuming \"Stable\"")
		s.userBranch = "Stable"
	}

	err = s.LoadConfig()
	if err != nil {
		fmt.Println("No config found. Trying to create default config file")
		s.autostartChanged = true
	}

	v, err := strconv.Atoi(strings.Replace(s.Version, ".", "", -1))
	if err != nil {
		v = 0
	}
	// add additional categories with new version
	if v < 112 {
		s.AvailableCategories = availableCategories
	}

	s.Version = version

	err = s.createDesktopFile()
	if err != nil {
		print(err)
	}

	err = s.SaveConfig(false)
	if err != nil {
		return nil, err
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

	err = s.createDesktopFile()
	if err != nil {
		return err
	}
	return nil
}

func (s *Config) createDesktopFile() error {
	asdir := s.configBaseDir + "/autostart"
	fileName := asdir + "/mntray.desktop"

	// make sure we have an autostart dir
	_, err := os.Stat(asdir)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		err = os.Mkdir(asdir, 0755)
		if err != nil {
			return err
		}
	}

	// if autostart file is not present, make sure to create it
	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			s.autostartChanged = true
		}
	}

	// if autostart setting has changed, recreate the .desktop file
	if s.autostartChanged {
		// delete previous desktop file
		os.Remove(fileName)

		_, err = os.Stat(fileName)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}

			var fb []byte

			if s.Autostart {
				fb = []byte(desktopFile)
			} else {
				fb = []byte(desktopFile + noAuto)
			}

			err = ioutil.WriteFile(fileName, fb, 0644)
			if err != nil {
				print(err)
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

	filteredArticles := []Article{}
	var catList []string

	if s.SetCategoriesFromBranch {
		arch := ""
		if runtime.GOARCH == "arm64" {
			arch = "ARM "
		}
		catList = append(s.AddCategoriesBranch, arch+s.userBranch+" Updates")
	} else {
		catList = s.Categories
	}

	for _, a := range *Articles {
		for _, c := range catList {
			if c == a.Category {
				filteredArticles = append(filteredArticles, a)
			}
		}
	}

	sort.Slice(filteredArticles, func(i, j int) bool {
		return filteredArticles[i].PublishedDate.Unix() < filteredArticles[j].PublishedDate.Unix()
	})

	return filteredArticles, nil
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
				return strings.Replace(strings.Replace(strings.Title(b[1]), "Arm-", "", -1), "-", " ", -1)
			}
			return ""
		}
	}
	return ""

}
