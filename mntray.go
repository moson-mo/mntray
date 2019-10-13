package main

import (
	"fmt"

	"github.com/moson-mo/mntray/tray"
)

func main() {
	err := tray.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
