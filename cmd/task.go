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
		util.Log.Debug("开启p站同步")
		pixivCycleTask()
	}
}

func pixivCycleTask() {
	util.Log.Info("start pixiv cycle task")
	t, err := tgbot.InitBot()
	if err != nil {
		util.Log.Errorf("init bot error: %v", err)
	}

	for {
		newUrls, err := pixiv.NewFavURLs()
		if err != nil {
			util.Log.Errorf("get new fav urls error: %v", err)
		}
		err = t.SendPhotosToChan(t, config.Conf.Storages.TelegramChannel.UserName, newUrls)
		if err != nil {
			util.Log.Errorf("send photos to channel error: %v", err)
		}
		util.Log.Infof("success,sleep %d minutes", config.Conf.Sources.Pixiv.RefreshTime)
		time.Sleep(time.Duration(config.Conf.Sources.Pixiv.RefreshTime) * time.Minute)
	}
}
