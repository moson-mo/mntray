package tray

import (
	"os"
	"os/exec"
	"sort"
	"sync"
	"time"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/gui"
	ui "github.com/therecipe/qt/widgets"
)

type SystemTrayIcon struct {
	ui.QSystemTrayIcon

	_ func(a article, b bool) `slot:"triggerSlot"`
	_ func(err error)         `slot:"connectionDead"`
	_ func()                  `slot:"hideIcon"`
}

type menuItem struct {
	item *ui.QAction
	news article
}

type trayIcon struct {
	app          *ui.QApplication
	icon         *SystemTrayIcon
	newsMenu     *ui.QMenu
	config       *settings
	items        []menuItem
	cacheDir     string
	quit         chan bool
	mut          *sync.Mutex
	ico          *gui.QIcon
	icoNew       *gui.QIcon
	icoChecked   *gui.QIcon
	icoUnchecked *gui.QIcon
	icoExit      *gui.QIcon
	lastMenuItem *menuItem
	wg           *sync.WaitGroup
}

func Run() error {
	app := ui.NewQApplication(len(os.Args), os.Args)

	// load config from file, if not existing create new config file with defaults
	config, err := NewSettings()
	if err != nil {
		return err
	}

	// initialize the tray icon
	ti := newIcon(app, config)
	ti.icon.Show()

	// load articles from saved file
	arts, err := config.LoadArticles()
	if err != nil {
		print("Error loading articles from file")
	}

	// create menu items for loaded articles
	for _, a := range arts {
		ti.receivedNews(a, true)
	}

	// initialize websocket and wait for news
	ti.waitForNews()

	// start main loop
	ui.QApplication_Exec()

	// wait for go-routine to finish
	ti.wg.Wait()
	return nil
}

// creates new tray icon and initializes variables
func newIcon(app *ui.QApplication, conf *settings) *trayIcon {
	ti := trayIcon{
		app:          app,
		icon:         NewSystemTrayIcon(app),
		quit:         make(chan bool, 1),
		config:       conf,
		mut:          &sync.Mutex{},
		ico:          gui.QIcon_FromTheme2("mntray-regular", gui.NewQIcon5(":assets/images/mntray-regular.png")),
		icoNew:       gui.QIcon_FromTheme2("mntray-news", gui.NewQIcon5(":assets/images/mntray-news.png")),
		icoChecked:   gui.QIcon_FromTheme2("emblem-checked", gui.QIcon_FromTheme2("emblem-default", gui.QIcon_FromTheme("dialog-yes"))),
		icoExit:      gui.QIcon_FromTheme("application-exit"),
		icoUnchecked: gui.QIcon_FromTheme2("vcs-removed", gui.QIcon_FromTheme("emblem-draft")),
		wg:           &sync.WaitGroup{},
	}

	ti.icon.SetIcon(ti.ico)
	ti.createNewsMenu()

	ti.icon.ConnectTriggerSlot(func(a article, fromFile bool) {
		// add menu item
		i := ti.addMenuItem(a)
		mi := menuItem{
			item: i,
			news: a,
		}
		ti.items = append(ti.items, mi)

		// set trayicon
		ti.setTrayIcon()

		// cleanup old menu items
		ti.cleanupMenu()

		// show bubble for unread articles
		if !fromFile && !a.WasRead && ti.containsArticle(a) {
			ti.lastMenuItem = &mi
			ti.icon.ShowMessage2("Manjaro News - New article", a.Title, ti.icoNew, 5000)
		}
	})

	// message when articles can not be downloaded (run on main thread)
	ti.icon.ConnectConnectionDead(func(err error) {
		ti.lastMenuItem = nil
		if ti.config.HideNoNews {
			ti.icon.Show()
		}
		ti.icon.ShowMessage2("Manjaro News - Error", "Error fetching news:\n"+err.Error(), gui.QIcon_FromTheme("error"), 5000)
		if ti.config.HideNoNews {
			go func() {
				<-time.After(5 * time.Second)
				ti.icon.HideIcon()
			}()
		}
	})

	// hide icon (run on main thread)
	ti.icon.ConnectHideIcon(func() {
		ti.icon.Hide()
	})

	// Open article when a message has been clicked (run on main thread)
	ti.icon.ConnectMessageClicked(func() {
		if ti.lastMenuItem != nil {
			ti.openArticle(ti.lastMenuItem.news)
			ti.lastMenuItem.item.SetIcon(ti.icoChecked)
		}
	})

	return &ti
}

// creates the news menu
func (ti *trayIcon) createNewsMenu() {
	menu := ui.NewQMenu2("news-menu", nil)
	ti.icon.ConnectActivated(func(reason ui.QSystemTrayIcon__ActivationReason) {
		if reason == ui.QSystemTrayIcon__Trigger {
			ti.newsMenu.Exec2(gui.QCursor_Pos(), nil)
		}
	})
	menu.AddSeparator()
	mr := menu.AddAction("Mark all as read")
	mr.ConnectTriggered(func(bool) {
		ti.markAllRead()
		ti.save()
	})

	quit := menu.AddAction("Quit")
	quit.SetIcon(ti.icoExit)
	quit.ConnectTriggered(func(bool) {
		ti.saveAndQuit()
	})
	quit.SetData(core.NewQVariant7(0))
	ti.newsMenu = menu
	ti.icon.SetContextMenu(menu)
}

// save articles and settings to disk
func (ti *trayIcon) save() {
	articles := make([]article, len(ti.items))
	for i := range ti.items {
		articles[i] = ti.items[i].news
	}
	ti.config.SaveArticles(articles)
	err := ti.config.SaveSettings(true)
	if err != nil {
		print("Could not save settings: " + err.Error())
	}
}

// saves articles to disk and quits the application
func (ti *trayIcon) saveAndQuit() {
	ti.save()
	ti.quit <- true
	ti.app.Quit()
}

// adds an item to the news menu
func (ti *trayIcon) addMenuItem(a article) *ui.QAction {
	var item *ui.QAction

	if len(ti.newsMenu.Actions()) == 0 {
		item = ti.newsMenu.AddAction(a.Title)
	} else {
		item = ui.NewQAction2(a.Title, ti.newsMenu)

		ti.newsMenu.InsertAction(ti.getInsertPosition(a), item)
	}
	item.SetData(core.NewQVariant7(a.PublishedDate.Unix()))

	item.SetIcon(ti.icoUnchecked)
	if a.WasRead {
		item.SetIcon(ti.icoChecked)
	}

	item.ConnectTriggered(func(checked bool) {
		ti.openArticle(a)
	})
	return item
}

// opens the article in the default browser
func (ti *trayIcon) openArticle(a article) {
	ti.markRead(a)
	print("Opened item: " + a.Title)

	com := exec.Command("xdg-open", a.URL)
	com.Start()
	ti.save()
}

// called when we got news from the server
func (ti *trayIcon) receivedNews(a article, fromFile bool) {
	if !fromFile && ti.containsArticle(a) {
		return
	}

	// add new item to menu, we use a slot to run on the main thread.
	ti.icon.TriggerSlot(a, fromFile)
}

// checks if we already have this article in our menu
func (ti *trayIcon) containsArticle(a article) bool {
	for _, i := range ti.items {
		if i.news.GUID == a.GUID {
			return true
		}
	}
	return false
}

// returns the position where the menu item is to be inserted (based on the date)
func (ti *trayIcon) getInsertPosition(a article) *ui.QAction {
	actions := ti.newsMenu.Actions()
	pos := actions[0]
	for _, i := range actions {
		if a.PublishedDate.Unix() > i.Data().ToLongLong(nil) {
			pos = i
			break
		}
	}
	return pos
}

// marks an article as read and updates the trayicon
func (ti *trayIcon) markRead(a article) {
	for i := range ti.items {
		if ti.items[i].news.GUID == a.GUID {
			ti.items[i].news.WasRead = true
			ti.items[i].item.SetIcon(ti.icoChecked)
		}
	}
	ti.setTrayIcon()
}

// marks all articles as read and updates the trayicon
func (ti *trayIcon) markAllRead() {
	for i := range ti.items {
		ti.items[i].news.WasRead = true
		ti.items[i].item.SetIcon(ti.icoChecked)
	}
	ti.setTrayIcon()
}

// remove items if exceeding maxitems
func (ti *trayIcon) cleanupMenu() {
	sort.Slice(ti.items, func(i, j int) bool {
		return ti.items[i].news.PublishedDate.Unix() < ti.items[j].news.PublishedDate.Unix()
	})

	x := len(ti.items) - ti.config.MaxArticles
	if x > 0 {
		rem := ti.items[0:x]
		for i := 0; i < len(rem); i++ {
			rem[i].item.DestroyQAction()
		}
		ti.items = ti.items[x:]
	}
}

// sets tray icon according to status (read / unread items)
func (ti *trayIcon) setTrayIcon() {
	for _, i := range ti.items {
		if !i.news.WasRead {
			if ti.config.HideNoNews {
				ti.icon.Show()
			}
			ti.icon.SetIcon(ti.icoNew)
			return
		}
	}
	if ti.config.HideNoNews {
		ti.icon.Hide()
	}
	ti.icon.SetIcon(ti.ico)
}
