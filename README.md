<img src="https://raw.githubusercontent.com/moson-mo/mntray/master/assets/images/mntray.png?inline=true"  align="left" width="120" />

# mntray
## A Manjaro Linux announcements notification app
</br>

[![Build Status](https://travis-ci.org/moson-mo/mntray.svg?branch=master)](https://travis-ci.org/moson-mo/mntray)
[![Go Report Card](https://goreportcard.com/badge/github.com/moson-mo/mntray)](https://goreportcard.com/report/github.com/moson-mo/mntray)
[![GitHub release](https://img.shields.io/github/v/tag/moson-mo/mntray.svg?label=release&sort=semver)](https://github.com/moson-mo/mntray/releases)

A small app which informs about announcements from manjaro.\
It creates a tray icon with a menu showing the latest announcements from the Manjaro forum RSS feed.

Announcements are retrieved from a http server (see [mnserver](https://github.com/moson-mo/mnserver/)) via post request.

This project is based on [Qt](https://www.qt.io) and the [Qt binding package](https://github.com/therecipe/qt) for golang.\
In order to run this app the `qt5-base` package needs to be installed on your system.

Why is it connecting to a server application rather then parsing the RSS feed directly?\
The RSS feed can be quite large (around 300 to 500 KB).\
Instead of downloading this file from the Manjaro forums host on a regular basis, it fetches news from mnserver.\
There's much less data to be transferred and less burden on the forum host and client since the data is stripped down to the bare minimum.
</br>

## How to install

Binaries are available from the [releases](https://github.com/moson-mo/mntray/releases) page.
</br>

## How to build

* Install go from your package manager or download it from the [Golang](https://golang.org/dl/) site. 
* Install Qt bindings for Golang: `go get -u -v github.com/therecipe/qt/cmd/...`
* Download this package with: `go get -d github.com/moson-mo/mntray`
* Change to package dir: `cd $(go env GOPATH)/src/github.com/moson-mo/mntray/`
* Set environment variable QT_PKG_CONFIG: `export QT_PKG_CONFIG=true`
* Run: `$(go env GOPATH)/bin/qtdeploy build desktop` from the mntray directory
* The binary will be in the `deploy/linux/` dir

For further information & cross compilation options please have a look at [this wiki](https://github.com/therecipe/qt/wiki)

#### Build with docker image

* Install docker with your package manager
* Add your user account to the docker group or run the following commands as root
* Download docker image `docker pull therecipe/qt:linux` / `sudo docker pull therecipe/qt:linux`
* Build with `$(go env GOPATH)/bin/qtdeploy -docker build linux` / `sudo -E $(go env GOPATH)/bin/qtdeploy -docker build linux`
</br>

## Configuration

On the first startup a config file (`~/.config/mntray/settings.json`) is created with some default settings.
You can either use the GUI to change the configuration (open "Settings" from the menu) or edit the config file.

```
{
	"Version": "0.2.0"
	"ServerURL": "http://manjaro.moson.eu:10111/news",
	"MaxArticles": 10,
	"AvailableCategories": [
		"Testing Updates",
		"Stable Updates",
		"Unstable Updates",
		"Announcements",
		"manjaro32",
		"Twitter"
	],
	"Categories": [
		"Testing Updates",
		"Stable Updates",
		"Unstable Updates",
		"Announcements",
		"manjaro32"
	],
	"AddCategoriesBranch": [
		"Announcements",
		"Twitter"
	],
	"RefreshInterval": 600,
	"HideNoNews": false,
	"Autostart": true,
	"ErrorNotifications": true,
	"DelayAfterStart": 60,
	"SetCategoriesFromBranch": true
}
```

Option | Description
--- | ---
Version| Version number. Do not change!|
URL| WebSocket URL of the mnservice server|
MaxArticles| The maximum number of articles to retrieve / show in the menu|
AvailableCategories| The categories that available for subscription. Do not change!|
Categories| The categories you want to get announcements for</br>Remove unwanted categories if needed</br></br>**note:* Is ignored when SetCategoriesFromBranch is "true"|
AddCategoriesBranch| The categories you want to get announcements for</br>additional to the branch you are using</br></br>**note:* Is ignored when SetCategoriesFromBranch is "false"|
RefreshInterval| The interval (in seconds) in which mntray will check for new articles|
Autostart| Places a .desktop file in the users autostart folder when "true"|
HideNoNews| When set to "true", the tray icon is hidden when all news have been read</br></br>**note:* Does not work reliably on GNOME & KDE. See "Known issues"|
ErrorNotifications| Show a notification in case articles can not be retrieved (f.e. network down)|
DelayAfterStart| Delays checking for news articles after startup (in seconds), f.e. wait for network to be up.</br></br> **note:* This setting only takes effect when mntray is started with parameter "--delay"|
SetCategoriesFromBranch| If "true", it auto-detects the Manjaro branch and filters categories accordingly (f.e. "Stable Updates" & "Announcements")</br></br>|

</br>

## Dependencies

* [qt](https://github.com/therecipe/qt) - Qt binding for Go
</br>

## Screenshots

#### Tray icon / menu

![xfce menu](https://github.com/moson-mo/mntray/raw/master/screenshots/xfce_menu.png?inline=true)
![kde menu](https://github.com/moson-mo/mntray/raw/master/screenshots/kde_menu.png?inline=true)
![gnome menu](https://github.com/moson-mo/mntray/raw/master/screenshots/gnome_menu.png?inline=true)

#### Notifications

![xfce notification](https://github.com/moson-mo/mntray/raw/master/screenshots/xfce_notification.png?inline=true)
![kde notification](https://github.com/moson-mo/mntray/raw/master/screenshots/kde_notification.png?inline=true)
![gnome notification](https://github.com/moson-mo/mntray/raw/master/screenshots/gnome_notification.png?inline=true)
</br>

#### Settings dialog

![xfce settings](https://github.com/moson-mo/mntray/raw/master/screenshots/xfce_settings.png?inline=true)
![kde settings](https://github.com/moson-mo/mntray/raw/master/screenshots/kde_settings.png?inline=true)
![gnome settings](https://github.com/moson-mo/mntray/raw/master/screenshots/gnome_settings.png?inline=true)
</br>

## Thanks to

* The Manjaro Team for the great distro
* The Manjaro community for testing and feedback
* [bogdancovaciu](https://github.com/bogdancovaciu) for providing the icons
* [SGS](https://github.com/sgse) for providing the app logo
* [therecipe](https://github.com/therecipe) for the Qt binding library
* The ones I forgot to mention here :)
</br>

## Todo / Plans

* Write tests
</br>

## Known issues

* "HideNoNews" not working correctly\
\
There seems to be a problem with the Hide() method of QSystemTrayIcon (Qt).\
Due to that this function does not work reliably on GNOME and KDE environments atm.