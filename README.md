<img src="https://raw.githubusercontent.com/moson-mo/mntray/master/assets/images/mntray.png?inline=true"  align="left" width="120" />

# mntray
## A Manjaro Linux announcements notification app
</br>

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
When you want to chang settings, make sure the app is not running since it will save settings on exit and overwrite your changes.

```
{
	"ServerURL": "http://manjaro.moson.eu:10111/news",
	"MaxArticles": 10,
	"Categories": [
		"Testing Updates",
		"Stable Updates",
		"Unstable Updates",
		"Announcements",
		"manjaro32"
	],
	"ArticlesFile": "/home/moson/.config/mntray/articles.json",
	"RefreshInterval": 600,
	"HideNoNews": false,
	"Autostart": true,
	"ErrorNotifications": true
}
```

Option | Description
--- | ---
URL| WebSocket URL of the mnservice server|
MaxArticles| The maximum number of articles to retrieve / show in the menu|
Categories| The categories you want to get announcements for</br>Remove unwanted categories if needed|
ArticlesFile| Path to the local file with news articles|
RefreshInterval| The interval (in seconds) in which mntray will check for new articles|
Autostart| Places a .desktop file in the users autostart folder when "true"|
HideNoNews| When set to "true", the tray icon is hidden when all news have been read|
ErrorNotifications| Show a notification in case articles can not be retrieved (f.e. network down)|
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

## Thanks to

* The Manjaro Team for the great distro
* The Manjaro community for testing and feedback
* [bogdancovaciu](https://github.com/bogdancovaciu) for providing the icons
* [SGS](https://github.com/sgse) for providing the app logo
* [therecipe](https://github.com/therecipe) for the Qt binding library
* The ones I forgot to mention here :)
</br>

## Todo / Plans

* Code cleanup (things are getting messy)
* Unit tests
* Integrate twitter news
* GUI settings editor
* Set initial category according to user's branch
* Update the screenshots 😜
