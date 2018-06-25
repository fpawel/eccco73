package main

import (
	"os"
	"github.com/lxn/win"
	"syscall"
	"path/filepath"
	"github.com/lxn/walk"
	"fmt"
)

const (
	appName = "eccco73"

)

func appFolderPath() string {
	var appDataPath string
	if appDataPath = os.Getenv("MYAPPDATA"); len(appDataPath) == 0 {
		var buf [win.MAX_PATH]uint16
		if !win.SHGetSpecialFolderPath(0, &buf[0], win.CSIDL_APPDATA, false) {
			panic("SHGetSpecialFolderPath failed")
		}
		appDataPath = syscall.UTF16ToString(buf[0:])
	}
	appDataPath = filepath.Join(appDataPath, "Аналитприбор", appName)
	_,err := os.Stat(appDataPath)
	if err != nil{
		if os.IsNotExist(err) { // создать каталог если его нет
			os.Mkdir(appDataPath, os.ModePerm)
		} else {
			panic(err)
		}
	}
	return appDataPath
}

func appFolderFileName(filename string) string {
	return filepath.Join(appFolderPath(), filename)
}

func  appConfigFileName() string{
	return filepath.Join(appFolderFileName("app.config"))
}


func setupWalkApplication(app *walk.Application) {

	app.SetOrganizationName("Аналитприбор")
	app.SetProductName(appName)
	appFolderPath() // создать каталог с данными и настройками программы если его нет
	sets := walk.NewIniFileSettings("settings.ini")
	if err := sets.Load(); err != nil {
		fmt.Println("load settings.ini error:", err)
	}
	app.SetSettings(sets)
}

