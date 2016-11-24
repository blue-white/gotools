package golog_test

import (
	"github.com/blue-white/gotools/golog"
	"testing"
)

func TestGolog(t *testing.T) {
	debug := golog.NewLogger(golog.Debug)
	debug.Debug("debug level: debug")
	debug.Info("debug level: info")
	debug.Error("debug level: error")

	info := golog.NewLogger(golog.Info)
	info.Debug("info level: debug")
	info.Info("info level: info")
	info.Error("info level: error")

	err := golog.NewLogger(golog.Error)
	err.Debug("err level: debug")
	err.Info("err level: info")
	err.Error("err level: error")

	fatal := golog.NewLogger(golog.Fatal)
	fatal.Debug("err level: debug")
	fatal.Info("err level: info")
	fatal.Error("err level: error")

	off := golog.NewLogger(golog.OFF)
	off.Debug("off level: debug")
	off.Info("off level: info")
	off.Error("off level: error")
	off.Fatal("off level: fatal")
}
