package util

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

func newlog(LogChannelName string) *slog.Logger {
	defer slog.MustFlush()
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
	slog.SetLogLevel(slog.InfoLevel)
	slog.DefaultChannelName = LogChannelName
	return slog.Std().Logger
}

var Log = newlog("favpics-helper")
