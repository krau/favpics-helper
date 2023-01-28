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
		pixivToTgChanTask()
	}
}

func pixivToTgChanTask() {
	util.Log.Info("start pixiv to tg chan task")
	for {
		util.Log.Debug("start getting new fav pics")
		pics, err := pixiv.Pixiv{}.NewFavPics()
		if err != nil {
			util.Log.Error("get new fav pics error:", err)
		}
		if len(pics) > 0 {
			util.Log.Debug("get new fav pics done")
			for _, pic := range pics {
				err = tgbot.SendPicsToChan(config.Conf.Storages.TelegramChannel.UserName, pic)
				if err != nil {
					util.Log.Error("send pic to tg chan error:", err)
					continue
				}
				util.Log.Debug("send pic done:", pic.Link)
			}
			util.Log.Infof("done,sleep %d minutes", config.Conf.Sources.Pixiv.RefreshTime)
			time.Sleep(time.Duration(config.Conf.Sources.Pixiv.RefreshTime) * time.Minute)
		} else {
			util.Log.Infof("no new fav pics,sleep %d minute", config.Conf.Sources.Pixiv.Interval)
			time.Sleep(time.Duration(config.Conf.Sources.Pixiv.Interval) * time.Minute)
		}
	}
}
