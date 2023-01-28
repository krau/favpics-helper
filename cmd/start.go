package cmd

import (
	"time"

	"github.com/krau/favpics-helper/internal/middlewares/tgbot"
	"github.com/krau/favpics-helper/internal/sources"
	"github.com/krau/favpics-helper/pkg/config"
	"github.com/krau/favpics-helper/pkg/util"
)

func Start() {
	util.Log.Notice("start favpics-helper")
	pixivTricker := time.NewTicker(time.Duration(config.Conf.Sources.Pixiv.Interval) * time.Minute)
	twitterTricker := time.NewTicker(time.Duration(config.Conf.Sources.Twitter.Interval) * time.Minute)
	go twitterToTgChanTask()
	time.Sleep(time.Second * 5)
	go pixivToTgChanTask()
	for {
		select {
		case <-pixivTricker.C:
			go pixivToTgChanTask()
		case <-twitterTricker.C:
			go twitterToTgChanTask()
		}
	}
}

func pixivToTgChanTask() {
	util.Log.Info("start pixiv to tg chan task")
	pixiv := sources.Pixiv{}
	util.Log.Debug("start getting new fav pics")
	pics, err := pixiv.NewFavPics()
	if err != nil {
		util.Log.Error("get new fav pics error:", err)
	}
	util.Log.Debug("get new fav pics done")
	for _, pic := range pics {
		err = tgbot.SendPicsToChan(config.Conf.Storages.TelegramChannel.UserName, pic)
		if err != nil {
			util.Log.Error("send pic to tg chan error:", err)
			continue
		}
		util.Log.Debug("send pic done:", pic.Link)
	}
}

func twitterToTgChanTask() {
	util.Log.Info("start twitter to tg chan task")
	twitter := sources.Twitter{}
	util.Log.Debug("start getting new fav pics")
	pics, err := twitter.NewFavPics()
	if err != nil {
		util.Log.Error("get new fav pics error:", err)
	}
	util.Log.Debug("get new fav pics done")
	for _, pic := range pics {
		err = tgbot.SendPicsToChan(config.Conf.Storages.TelegramChannel.UserName, pic)
		if err != nil {
			util.Log.Error("send pic to tg chan error:", err)
			continue
		}
		util.Log.Debug("send pic done:", pic.Link)
	}
}
