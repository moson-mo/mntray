package tray

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sort"
	"strconv"
	"syscall"
	"time"

	"github.com/therecipe/qt/uitools"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	ui "github.com/therecipe/qt/widgets"
)

// SystemTrayIcon is a customized QSystemTrayIcon
type SystemTrayIcon struct {
	ui.QSystemTrayIcon

	_ func(Article, bool) `slot:"newArticleSlot"`
	_ func(error)         `slot:"errorSlot"`
	_ func()              `slot:"hideIconSlot"`
}

var (
	ico, icoNew, icoChecked, icoExit, icoUnchecked *gui.QIcon
)

// TrayIcon is our main struct. It holds a list of articles, the tray icon and menu
type TrayIcon struct {
	App             *ui.QApplication
	Icon            *SystemTrayIcon
	Menu            *ui.QMenu
	SettingsDialog  *ui.QDialog
	SettingsWidgets *SettingsWidgets
	Conf            *Config
	Articles        []Article
	LastArticle     *Article
	Delay           bool
}

type SettingsWidgets struct {
	btnBox *ui.QDialogButtonBox

	txtURL             *ui.QLineEdit
	txtDelayStart      *ui.QLineEdit
	txtRefreshInterval *ui.QLineEdit
	txtMaxArticles     *ui.QLineEdit

	cbAutostart          *ui.QCheckBox
	cbErrorNotifications *ui.QCheckBox
	cbHideWhenRead       *ui.QCheckBox
	cbSetCatBranch       *ui.QCheckBox

	listCategories *ui.QListWidget

	lblCategories *ui.QLabel
}

// NewTrayIcon creates a new tray icon
func NewTrayIcon(delay bool) error {
	var err error
	t := &TrayIcon{
		App:             ui.NewQApplication(len(os.Args), os.Args),
		Delay:           delay,
		Articles:        []Article{},
		SettingsWidgets: &SettingsWidgets{},
	}

	// load settings window
	t.App.SetQuitOnLastWindowClosed(false)
	f := core.NewQFile2(":assets/settings.ui")
	l := uitools.NewQUiLoader(t.App)
	w := l.Load(f, nil)
	t.SettingsDialog = ui.NewQDialogFromPointer(w.Pointer())
	t.SettingsDialog.ConnectFinished(t.onDialogClosed)
	t.setSettingsWidgets()
	//t.setSettingsWidgetValues()

	// set icons
	ico = gui.QIcon_FromTheme2("mntray-regular", gui.NewQIcon5(":assets/images/mntray-regular.png"))
	icoNew = gui.QIcon_FromTheme2("mntray-news", gui.NewQIcon5(":assets/images/mntray-news.png"))
	icoChecked = gui.QIcon_FromTheme2("emblem-checked", gui.QIcon_FromTheme2("emblem-default", gui.QIcon_FromTheme("dialog-yes")))
	icoExit = gui.QIcon_FromTheme("application-exit")
	icoUnchecked = gui.QIcon_FromTheme2("vcs-removed", gui.QIcon_FromTheme("emblem-draft"))

	// create icon
	t.Icon = NewSystemTrayIcon(t.App)
	t.Icon.SetIcon(ico)
	t.Conf, err = NewConfig()
	if err != nil {
		return err
	}

	// load articles from file
	arts, err := t.Conf.LoadArticles()
	if err != nil {
		fmt.Println("Error loading articles from file:", err)
	}

	// make menu
	t.createMenu()

	// connect slots
	t.Icon.ConnectNewArticleSlot(t.onNewArticle)
	t.Icon.ConnectErrorSlot(t.onError)
	t.Icon.ConnectHideIconSlot(t.onHideIcon)

	// connect signals
	t.Icon.ConnectMessageClicked(t.onMessageClicked)

	// create menu items for loaded articles
	for i := range arts {
		t.gotNewArticle(arts[i], true)
	}

	// show icon
	t.Icon.Show()
	t.setTrayIcon()

	// start loop for fetching news
	t.waitForNews()

	// handle signals from OS
	go t.handleOSSignals()

	// Qt main loop
	t.App.Exec()
	return nil
}

// creates the basic menu
func (t *TrayIcon) createMenu() {
	m := ui.NewQMenu(nil)
	t.Icon.ConnectActivated(t.onActivated)

	m.AddSeparator()
	mr := m.AddAction("Mark all as read")
	mr.ConnectTriggered(t.onMarkAsReadClicked)

	set := m.AddAction("Settings")
	set.SetIcon(gui.QIcon_FromTheme("settings"))
	set.ConnectTriggered(t.onSettingsClicked)

	quit := m.AddAction("Quit")
	quit.SetIcon(icoExit)
	quit.ConnectTriggered(t.onQuitClicked)

	t.Menu = m
	t.Icon.SetContextMenu(m)
}

// assign ui widgets
func (t *TrayIcon) setSettingsWidgets() {
	w := t.SettingsWidgets
	d := t.SettingsDialog
	w.txtURL = ui.NewQLineEditFromPointer(d.FindChild("txtURL", core.Qt__FindChildrenRecursively).Pointer())
	w.txtDelayStart = ui.NewQLineEditFromPointer(d.FindChild("txtDelayStart", core.Qt__FindChildrenRecursively).Pointer())
	w.txtMaxArticles = ui.NewQLineEditFromPointer(d.FindChild("txtMaxArticles", core.Qt__FindChildrenRecursively).Pointer())
	w.txtRefreshInterval = ui.NewQLineEditFromPointer(d.FindChild("txtRefreshInterval", core.Qt__FindChildrenRecursively).Pointer())
	w.cbAutostart = ui.NewQCheckBoxFromPointer(d.FindChild("cbAutostart", core.Qt__FindChildrenRecursively).Pointer())
	w.cbErrorNotifications = ui.NewQCheckBoxFromPointer(d.FindChild("cbErrorNotifications", core.Qt__FindChildrenRecursively).Pointer())
	w.cbHideWhenRead = ui.NewQCheckBoxFromPointer(d.FindChild("cbHideWhenRead", core.Qt__FindChildrenRecursively).Pointer())
	w.cbSetCatBranch = ui.NewQCheckBoxFromPointer(d.FindChild("cbSetCatBranch", core.Qt__FindChildrenRecursively).Pointer())
	w.listCategories = ui.NewQListWidgetFromPointer(d.FindChild("listCategories", core.Qt__FindChildrenRecursively).Pointer())
	w.btnBox = ui.NewQDialogButtonBoxFromPointer(d.FindChild("buttonBox", core.Qt__FindChildrenRecursively).Pointer())
	w.lblCategories = ui.NewQLabelFromPointer(d.FindChild("lblCategories", core.Qt__FindChildrenRecursively).Pointer())

	w.cbSetCatBranch.ConnectToggled(t.onSetCategoriesFromBranchToggled)
	w.btnBox.Button(ui.QDialogButtonBox__Reset).ConnectClicked(t.onResetButtonClicked)
}

// set ui widget values from config
func (t *TrayIcon) setSettingsWidgetValues() {
	w := t.SettingsWidgets
	c := t.Conf
	w.txtURL.SetText(c.ServerURL)
	w.txtDelayStart.SetText(strconv.Itoa(c.DelayAfterStart))
	w.txtMaxArticles.SetText(strconv.Itoa(c.MaxArticles))
	w.txtRefreshInterval.SetText(strconv.Itoa(c.RefreshInterval))

	w.cbAutostart.SetChecked(c.Autostart)
	w.cbErrorNotifications.SetChecked(c.ErrorNotifications)
	w.cbHideWhenRead.SetChecked(c.HideNoNews)
	w.cbSetCatBranch.SetChecked(c.SetCategoriesFromBranch)

	w.listCategories.Clear()
	w.listCategories.AddItems(c.AvailableCategories)

	t.setSettingsCategory(c.SetCategoriesFromBranch)
}

// set settings dialog categories from conf / based on checkbox
func (t *TrayIcon) setSettingsCategory(fromBranch bool) {
	lc := t.SettingsWidgets.listCategories
	for i := 0; i < lc.Count(); i++ {
		lc.Item(i).SetSelected(false)
		lc.Item(i).SetHidden(false)
	}

	if !fromBranch {
		for _, c := range t.Conf.Categories {
			items := lc.FindItems(c, core.Qt__MatchExactly)
			if len(items) == 1 {
				items[0].SetSelected(true)
			}
		}
	} else {
		items := lc.FindItems(" Updates", core.Qt__MatchContains)
		for _, i := range items {
			i.SetHidden(true)
		}
		for _, c := range append(t.Conf.AddCategoriesBranch, t.Conf.userBranch+" Updates") {
			items := lc.FindItems(c, core.Qt__MatchExactly)
			if len(items) == 1 {
				items[0].SetSelected(true)
			}
		}
	}
}

// set conf from widget values
func (t *TrayIcon) setConfFromWidgets() {
	c := t.Conf
	w := t.SettingsWidgets
	c.ServerURL = w.txtURL.Text()
	c.RefreshInterval, _ = strconv.Atoi(w.txtRefreshInterval.Text())
	c.MaxArticles, _ = strconv.Atoi(w.txtMaxArticles.Text())
	c.DelayAfterStart, _ = strconv.Atoi(w.txtDelayStart.Text())

	c.Autostart = w.cbAutostart.IsChecked()
	c.ErrorNotifications = w.cbErrorNotifications.IsChecked()
	c.HideNoNews = w.cbHideWhenRead.IsChecked()
	c.SetCategoriesFromBranch = w.cbSetCatBranch.IsChecked()

	cats := []string{}
	for i := 0; i < w.listCategories.Count(); i++ {
		if w.listCategories.Item(i).IsSelected() {
			cats = append(cats, w.listCategories.Item(i).Text())
		}
	}
	if c.SetCategoriesFromBranch {
		c.AddCategoriesBranch = cats
	} else {
		c.Categories = cats
	}
}

// executes when reset button was clicked
func (t *TrayIcon) onResetButtonClicked(c bool) {
	t.setSettingsWidgetValues()
}

// executes when set categories from branch checkbox is toggled
func (t *TrayIcon) onSetCategoriesFromBranchToggled(checked bool) {
	t.setSettingsCategory(checked)
	if checked {
		t.SettingsWidgets.lblCategories.SetText("Additional categories")
	} else {
		t.SettingsWidgets.lblCategories.SetText("Selected categories")
	}
}

// executed when Settings dialog is closed
func (t *TrayIcon) onDialogClosed(code int) {
	if code == 1 {
		t.setConfFromWidgets()
		t.save(false)
	}
}

// shows menu on left click
func (t *TrayIcon) onActivated(r ui.QSystemTrayIcon__ActivationReason) {
	if r == ui.QSystemTrayIcon__Trigger {
		t.Menu.Exec2(gui.QCursor_Pos(), nil)
	}
}

// executes when "Settings" is clicked
func (t *TrayIcon) onSettingsClicked(b bool) {
	t.SettingsDialog.Show()
	t.setSettingsWidgetValues()
}

// executes when "Mark all as read" is clicked
func (t *TrayIcon) onMarkAsReadClicked(c bool) {
	t.markAllRead()
	t.save(true)
	t.setTrayIcon()
}

// executes when "Quit" is clicked
func (t *TrayIcon) onQuitClicked(c bool) {
	t.save(true)
	t.App.Quit()
}

//
func (t *TrayIcon) onHideIcon() {
	t.Icon.Hide()
}

// saves articles and config to disk
func (t *TrayIcon) save(loadBeforeSave bool) {
	err := t.Conf.SaveArticles(t.Articles)
	if err != nil {
		fmt.Println("Could not save articles to disk:", err)
	}
	err = t.Conf.SaveConfig(loadBeforeSave)
	if err != nil {
		fmt.Println("Could not save settings to disk:", err)
	}
}

// marks all articles as read
func (t *TrayIcon) markAllRead() {
	for i := range t.Articles {
		t.Articles[i].WasRead = true
		t.Articles[i].mi.SetIcon(icoChecked)
	}
}

// marks an article as read
func (t *TrayIcon) markRead(a Article) {
	for i := range t.Articles {
		if t.Articles[i].GUID == a.GUID {
			t.Articles[i].WasRead = true
			t.Articles[i].mi.SetIcon(icoChecked)
		}
	}
}

// sets the icon according to the current status (read / unread items)
func (t *TrayIcon) setTrayIcon() {
	for i := range t.Articles {
		if !t.Articles[i].WasRead {
			if t.Conf.HideNoNews {
				t.Icon.Show()
			}
			t.Icon.SetIcon(icoNew)
			return
		}
	}
	if t.Conf.HideNoNews {
		t.Icon.Hide()
	}
	t.Icon.SetIcon(ico)
}

// remove items if exceeding maxitems
func (t *TrayIcon) cleanupMenu() {
	sort.Slice(t.Articles, func(i, j int) bool {
		return t.Articles[i].PublishedDate.Unix() < t.Articles[j].PublishedDate.Unix()
	})

	x := len(t.Articles) - t.Conf.MaxArticles
	if x > 0 {
		rem := t.Articles[0:x]
		for i := 0; i < len(rem); i++ {
			rem[i].mi.DestroyQAction()
		}
		t.Articles = t.Articles[x:]
	}
}

// returns the position where the menu item is to be inserted (based on the date)
func (t *TrayIcon) getInsertPosition(a Article) *ui.QAction {
	actions := t.Menu.Actions()
	pos := actions[0]
	for _, i := range actions {
		if a.PublishedDate.Unix() > i.Data().ToLongLong(nil) {
			pos = i
			break
		}
	}
	return pos
}

// checks if we already have this article in our menu
func (t *TrayIcon) containsArticle(a Article) bool {
	for _, i := range t.Articles {
		if i.GUID == a.GUID {
			return true
		}
	}
	return false
}

// called when we got news from the server
func (t *TrayIcon) gotNewArticle(a Article, fromFile bool) {
	if !fromFile && t.containsArticle(a) {
		return
	}

	// add new item to menu, we use a slot to run on the main thread.
	t.Icon.NewArticleSlot(a, fromFile)
}

// opens the article in the default browser
func (t *TrayIcon) openArticle(a Article) {
	com := exec.Command("xdg-open", a.URL)
	com.Start()
	t.markRead(a)
	t.save(true)
	t.setTrayIcon()
}

// adds an item to the news menu
func (t *TrayIcon) addMenuItem(a Article) *ui.QAction {

	item := ui.NewQAction2(a.Title, t.Menu)
	t.Menu.InsertAction(t.getInsertPosition(a), item)
	item.SetData(core.NewQVariant7(a.PublishedDate.Unix()))

	if a.WasRead {
		item.SetIcon(icoChecked)
	} else {
		item.SetIcon(icoUnchecked)
	}

	item.ConnectTriggered(func(c bool) {
		t.openArticle(a)
	})

	return item
}

// when an article has been recieved
func (t *TrayIcon) onNewArticle(a Article, fromFile bool) {
	a.mi = t.addMenuItem(a)
	t.Articles = append(t.Articles, a)

	t.cleanupMenu()
	t.setTrayIcon()

	// show bubble for unread articles
	if !fromFile && !a.WasRead && t.containsArticle(a) {
		t.LastArticle = &a
		t.Icon.ShowMessage2("Manjaro News - New article", a.Title, icoNew, 5000)
	}
}

// when an error has occured
func (t *TrayIcon) onError(err error) {
	t.LastArticle = nil
	if t.Conf.HideNoNews {
		t.Icon.Show()
	}
	t.Icon.ShowMessage2("Manjaro News - Error", err.Error(), gui.QIcon_FromTheme("error"), 5000)
	if t.Conf.HideNoNews {
		go func() {
			<-time.After(5 * time.Second)
			t.Icon.HideIconSlot()
		}()
	}
}

// when a message has been clicked
func (t *TrayIcon) onMessageClicked() {
	if t.LastArticle != nil {
		t.openArticle(*t.LastArticle)
	}
}

// handles os signals for graceful shutdown (kill, interrupt, term)
func (t *TrayIcon) handleOSSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c
	fmt.Println("Got signal:", s)
	t.onQuitClicked(true)
}
