// +build !ubports

package main

/*
#cgo CFLAGS: -pipe -O2 -Wall -W -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -O2 -std=gnu++11 -Wall -W -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../moson-mo -I. -I/opt/Qt/5.13.0/gcc_64/include -I/opt/Qt/5.13.0/gcc_64/include/QtGui -I/opt/Qt/5.13.0/gcc_64/include/QtCore -I. -isystem /usr/include/libdrm -I/opt/Qt/5.13.0/gcc_64/mkspecs/linux-g++
#cgo LDFLAGS: -O1 -Wl,-rpath,/opt/Qt/5.13.0/gcc_64/lib
#cgo LDFLAGS:  /opt/Qt/5.13.0/gcc_64/lib/libQt5Gui.so /opt/Qt/5.13.0/gcc_64/lib/libQt5Core.so -lGL -lpthread
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
