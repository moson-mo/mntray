<img src="https://github.com/moson-mo/mntray/raw/master/screenshots/logo.png?inline=true"  align="left" width="120" />

# mntray
## A Manjaro Linux announcements notification app
</br>

A small app which informs about announcements from manjaro.\
It creates a tray icon with a menu showing the latest announcements from the Manjaro forum RSS feed.

Announcements are retrieved from a WebSocket server (see [mnservice](https://github.com/moson-mo/mnservice/)) 

This project is based on [Qt](https://www.qt.io) and the [Qt binding package](https://github.com/therecipe/qt) for golang.\
In order to run this app the `qt5-base` package needs to be installed on your system.

Why is it connecting to a server application rather then parsing the RSS feed directly?\
The RSS feed can be quite large (around 300 to 500 KB).\
Instead of downloading this file from the Manjaro forums host on a regular basis, it sets up a websocket connection and waits for the server to push the required data to the client.\
So there's much less data to be transferred and less burden on the forum host and client.
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
	"URL": "ws://manjaro.moson.eu:6666/articles",
	"MaxArticles": 10,
	"Categories": [
		"Testing Updates",
		"Stable Updates",
		"Unstable Updates",
		"Announcements",
		"manjaro32"
	],
	"ArticlesFile": "/home/moson/.config/mntray/articles.json",
	"ReconnectInterval": 10,
	"NumberOfRetries": 5
	"HideNoNews": false
}
```

Option | Description
--- | ---
URL| WebSocket URL of the mnservice server|
MaxArticles| The maximum number of articles to retrieve / show in the menu|
Categories| The categories you want to get announcements for</br>Remove unwanted categories if needed|
ArticlesFile| Path to the local file with news articles|
ReconnectInterval| If the connection to the server is lost, mntray will try to reconnect with this interval (in seconds)|
NumberOfRetries| Number of retries before mntray gives up trying to reconnect if the connection is lost|
HideNoNews| When set to "true", the tray icon is hidden when all news have been read|
</br>

## Dependencies

* [qt](https://github.com/therecipe/qt) - Qt binding for Go
* [websocket](https://github.com/gorilla/websocket) - WebSocket implementation for Go
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

## Thanks to

* The Manjaro Team for the great distro
* The Manjaro community for testing and feedback
* [bogdancovaciu](https://github.com/bogdancovaciu) for providing the icons
* [SGS](https://github.com/sgse) for providing the app logo
* [therecipe](https://github.com/therecipe) for the Qt binding library
* The ones I forgot to mention here :)