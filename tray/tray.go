package tray

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sort"
	"sync"
	"syscall"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/gui"
	ui "github.com/therecipe/qt/widgets"
)

type SystemTrayIcon struct {
	ui.QSystemTrayIcon

	_ func(a article, b bool) `slot:"triggerSlot"`
	_ func()                  `slot:"connectionDead"`
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
	done         chan bool
	mut          *sync.Mutex
	ico          *gui.QIcon
	icoNew       *gui.QIcon
	icoChecked   *gui.QIcon
	icoUnchecked *gui.QIcon
	icoExit      *gui.QIcon
	sigc         chan os.Signal
	lastMenuItem *menuItem
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
		fmt.Println("Error loading articles from file")
	}

	// create menu items for loaded articles
	for _, a := range arts {
		ti.receivedNews(a, true)
	}

	// initialize websocket and wait for news
	ti.waitForNews()

	// start main loop
	ui.QApplication_Exec()
	return nil
}

// creates new tray icon and initializes variables
func newIcon(app *ui.QApplication, conf *settings) *trayIcon {
	ti := trayIcon{
		app:          app,
		icon:         NewSystemTrayIcon(app),
		done:         make(chan bool),
		config:       conf,
		mut:          &sync.Mutex{},
		ico:          gui.NewQIcon5(":assets/images/m64.png"),
		icoNew:       gui.NewQIcon5(":assets/images/m64n.png"),
		icoChecked:   gui.QIcon_FromTheme2("emblem-checked", gui.QIcon_FromTheme2("emblem-default", gui.QIcon_FromTheme("dialog-yes"))),
		icoExit:      gui.QIcon_FromTheme("application-exit"),
		icoUnchecked: gui.QIcon_FromTheme2("vcs-removed", gui.QIcon_FromTheme("emblem-draft")),
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

	ti.icon.ConnectConnectionDead(func() {
		ti.tiredConnecting()
	})

	ti.icon.ConnectMessageClicked(func() {
		if ti.lastMenuItem != nil {
			ti.openArticle(ti.lastMenuItem.news)
			ti.lastMenuItem.item.SetIcon(ti.icoChecked)
		}
	})

	// check for os signals and save articles to disk when app is killed.
	ti.sigc = make(chan os.Signal, 1)
	signal.Notify(ti.sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-ti.sigc
		fmt.Println("Signal received: " + s.String())
		ti.saveAndQuit()
	}()

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
	quit := menu.AddAction("Quit")
	quit.SetIcon(ti.icoExit)
	quit.ConnectTriggered(func(bool) {
		ti.saveAndQuit()
	})
	quit.SetData(core.NewQVariant7(0))
	ti.newsMenu = menu
	ti.icon.SetContextMenu(menu)
}

// saves articles to disk and quits the application
func (ti *trayIcon) saveAndQuit() {
	articles := make([]article, len(ti.items))
	for i := range ti.items {
		articles[i] = ti.items[i].news
	}
	ti.config.SaveArticles(articles)
	err := ti.config.SaveSettings()
	if err != nil {
		fmt.Println("Could not save settings: " + err.Error())
	}
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
		if !checked {
			item.SetIcon(ti.icoChecked)
		}
		ti.openArticle(a)
	})
	return item
}

// opens the article in the default browser
func (ti *trayIcon) openArticle(a article) {
	ti.markRead(a)
	fmt.Println("Opened item: " + a.Title)

	com := exec.Command("xdg-open", a.URL)
	com.Start()
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

// marks an article as read and update the trayicon
func (ti *trayIcon) markRead(a article) {
	for i := range ti.items {
		if ti.items[i].news.GUID == a.GUID {
			ti.items[i].news.WasRead = true
		}
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

// connection could not be made, add item to let user reconnect
func (ti *trayIcon) tiredConnecting() {
	if ti.config.HideNoNews {
		ti.icon.Show()
	}
	ti.icon.ShowMessage2("Manjaro News - Connection lost", "Connection lost, try to reconnect manually.", gui.NewQIcon5(":assets/images/manjaro64.png"), 5000)
	a := ui.NewQAction2("Reconnect... (connection lost)", ti.newsMenu)
	a.SetIcon(gui.QIcon_FromTheme("view-refresh"))
	ti.newsMenu.InsertAction(ti.newsMenu.Actions()[len(ti.newsMenu.Actions())-1], a)
	a.ConnectTriggered(func(checked bool) {
		// try reconnecting and remove menu entry
		ti.waitForNews()
		a.DestroyQAction()
		if ti.config.HideNoNews {
			ti.icon.Hide()
		}
	})
}
