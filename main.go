package main

import (
	"github.com/krau/favpics-helper/cmd"
	"github.com/krau/favpics-helper/pkg/util"
)

func main() {
	cmd.CycleTask()
}

func testlog() {
	util.Log.Trace("this is a Trace log message")
	util.Log.Debug("this is a Debug log message")
	util.Log.Info("this is a Info log message")
	util.Log.Warn("this is a Warn log message")
	util.Log.Error("this is a Error log message")
	util.Log.Fatal("this is a Fatal log message")
	util.Log.Panic("this is a Panic log message")
}
