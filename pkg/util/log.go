package util

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/krau/favpics-helper/pkg/config"
)

func newlog(LogChannelName string) *slog.Logger {
	//defer slog.MustFlush()
	logLevel := config.Conf.System.Log.Level

	errorHandler := handler.MustRotateFile("./logs/error.log", rotatefile.EveryDay,
		handler.WithLogLevels(slog.DangerLevels),
		handler.WithCompress(true),
	)
	// NormalLevels 包含： slog.InfoLevel, slog.NoticeLevel, slog.DebugLevel, slog.TraceLevel
	normalHandler := handler.MustRotateFile("./logs/info.log", rotatefile.EveryDay,
		handler.WithLogLevels(slog.NormalLevels), handler.WithCompress(true))

	fullHandler := handler.MustRotateFile("./logs/full.log", rotatefile.EveryDay,
		handler.WithLogLevels(slog.AllLevels), handler.WithCompress(true))
	slog.PushHandlers(errorHandler, normalHandler, fullHandler)
	switch logLevel {
	case "debug":
		slog.SetLogLevel(slog.DebugLevel)
	case "info":
		slog.SetLogLevel(slog.InfoLevel)
	case "warn":
		slog.SetLogLevel(slog.WarnLevel)
	case "error":
		slog.SetLogLevel(slog.ErrorLevel)
	default:
		slog.SetLogLevel(slog.InfoLevel)
	}
	slog.DefaultChannelName = LogChannelName
	return slog.Std().Logger
}

var Log = newlog("favpics-helper")
