package tray

import (
	"fmt"
	"strconv"

	"github.com/therecipe/qt/core"
	ui "github.com/therecipe/qt/widgets"
)

// SettingsWidgets hold a reference to all widgets on the settings window
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

// assign ui widgets
func (t *TrayIcon) assignWidgets() {
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
func (t *TrayIcon) setWidgetValuesFromConf() {
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

	t.setWidgetCategory(c.SetCategoriesFromBranch)
}

// set settings dialog categories from conf / based on checkbox
func (t *TrayIcon) setWidgetCategory(fromBranch bool) {
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
		for _, c := range t.Conf.AddCategoriesBranch {
			items := lc.FindItems(c, core.Qt__MatchExactly)
			if len(items) == 1 {
				items[0].SetSelected(true)
			}
		}
	}
}

// set conf from widget values
func (t *TrayIcon) setConfFromWidgets() {
	restartRequired := false
	c := t.Conf
	w := t.SettingsWidgets
	c.ServerURL = w.txtURL.Text()
	c.RefreshInterval, _ = strconv.Atoi(w.txtRefreshInterval.Text())
	maxA, _ := strconv.Atoi(w.txtMaxArticles.Text())
	if maxA > c.MaxArticles {
		restartRequired = true
	}
	c.MaxArticles = maxA
	c.DelayAfterStart, _ = strconv.Atoi(w.txtDelayStart.Text())

	c.Autostart = w.cbAutostart.IsChecked()
	c.ErrorNotifications = w.cbErrorNotifications.IsChecked()
	c.HideNoNews = w.cbHideWhenRead.IsChecked()

	if c.SetCategoriesFromBranch != w.cbSetCatBranch.IsChecked() {
		restartRequired = true
		c.SetCategoriesFromBranch = w.cbSetCatBranch.IsChecked()
	}

	cats := []string{}
	for i := 0; i < w.listCategories.Count(); i++ {
		if w.listCategories.Item(i).IsSelected() {
			cats = append(cats, w.listCategories.Item(i).Text())
		}
	}
	if c.SetCategoriesFromBranch {
		if !stringSEqual(c.AddCategoriesBranch, cats) {
			c.AddCategoriesBranch = cats
			restartRequired = true
		}
	} else {
		if !stringSEqual(c.Categories, cats) {
			c.Categories = cats
			restartRequired = true
		}
	}
	if restartRequired {
		fmt.Println("restart required")
		t.Icon.ShowMessage2("Manjaro News - Information", "You've changed setting which require a restart of mntray. Please restart me.", ico, 5000)
	}
}

// executes when reset button was clicked
func (t *TrayIcon) onResetButtonClicked(c bool) {
	t.setWidgetValuesFromConf()
}

// executes when set categories from branch checkbox is toggled
func (t *TrayIcon) onSetCategoriesFromBranchToggled(checked bool) {
	t.setWidgetCategory(checked)
	if checked {
		t.SettingsWidgets.lblCategories.SetText("Additional categories")
	} else {
		t.SettingsWidgets.lblCategories.SetText("Selected categories")
	}
}

// executed when Settings dialog is closed
func (t *TrayIcon) onDialogAccepted() {
	t.setConfFromWidgets()
	t.saveConfig(false)
	t.cleanupMenu()
	t.setTrayIcon()
}
