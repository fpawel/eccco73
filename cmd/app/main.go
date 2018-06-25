package main

import (
		_ "runtime/cgo"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var buildTime = "(undefined)"
var majorVersion = 0
var minorVersion = 0
var bugFixVersion = 0

func main() {

	log.SetFlags(log.Lshortfile)

	//logrus.SetOutput(os.Stdout)
	//logrus.SetFormatter(&logrus.TextFormatter{})
	//eventLogHook,err := eventloghook.NewDefault("eccco73")
	//if err != nil {
	//	panic(err)
	//}
	//logrus.AddHook( eventLogHook )

	app := NewApp()

	app.mw.Run()

	app.Close()

}

