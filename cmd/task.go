package cmd

import (
	"time"

	"github.com/krau/favpics-helper/internal/middlewares/tgbot"
	"github.com/krau/favpics-helper/internal/sources/pixiv"
	"github.com/krau/favpics-helper/pkg/config"
	"github.com/krau/favpics-helper/pkg/util"
)

func CycleTask() {
	util.Log.Info("start cycle tasks")

	if config.Conf.Sources.Pixiv.Enabled {
		util.Log.Debug("pixiv enabled")
		pixivCycleTask()
	}
}

func pixivCycleTask() {
	util.Log.Info("start pixiv cycle task")
	for {
		util.Log.Debug("start getting new fav pics")
		pics, err := pixiv.Pixiv{}.NewFavPics()
		if err != nil {
			util.Log.Error("get new fav pics error:", err)
		}
		util.Log.Info("get new fav pics done")
		if len(pics) > 0 {
			util.Log.Info("start sending new fav pics")
			err = tgbot.SendPicsToChan(config.Conf.Storages.TelegramChannel.UserName, pics)
			if err != nil {
				util.Log.Error("send new fav pics error:", err)
			}
			util.Log.Infof("done,sleep %d minutes", config.Conf.Sources.Pixiv.RefreshTime)
			time.Sleep(time.Duration(config.Conf.Sources.Pixiv.RefreshTime) * time.Minute)
		} else {
			util.Log.Info("no new fav pics,sleep 5 minute")
			time.Sleep(5 * time.Minute)
		}
	}
}
