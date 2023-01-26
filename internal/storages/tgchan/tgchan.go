package tgchan

import "github.com/krau/favpics-helper/pkg/config"

type TelegramChannel struct {
	UserName string
}

var TgChan TelegramChannel = TelegramChannel{
	UserName: config.Conf.Storages.TelegramChannel.UserName,
}
