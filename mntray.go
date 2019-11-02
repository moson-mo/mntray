package main

import (
	"flag"
	"fmt"

	"github.com/moson-mo/mntray/tray"

	"github.com/marcsauter/single"
)

func main() {
	delay := flag.Bool("delay", false, "Delays fetching news articles by the duration (in seconds) specified in the config file")
	flag.Parse()

	single.Lockfile = "/tmp/mntray.lock"
	s := single.New("mntray")
	if err := s.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		fmt.Println("mntray already running, exiting")
		return
	} else if err != nil {
		fmt.Printf("Failed to acquire exclusive app lock: %v\n", err)
		return
	}
	defer s.TryUnlock()

	err := tray.NewTrayIcon(*delay)
	if err != nil {
		fmt.Println("Error setting up tray icon:", err)
		return
	}
}
