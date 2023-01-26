package tgbot

import (
	"time"

	"github.com/krau/favpics-helper/pkg/util"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/krau/favpics-helper/pkg/config"
)

type TelegramBot struct {
	Bot *tgbotapi.BotAPI
}

func InitBot() (t *TelegramBot, err error) {
	util.Log.Info("init bot")
	bot, err := tgbotapi.NewBotAPI(config.Conf.Middlewares.TelegramBot.Token)
	util.Log.Infof("Authorized on account %s", bot.Self.UserName)
	if err != nil {
		return nil, err
	}
	t = &TelegramBot{
		Bot: bot,
	}
	util.Log.Info("bot init success")
	return t, nil
}

func (t *TelegramBot) SendPhotosToChan(tgbot *TelegramBot, UserName string, urls []string) error {
	util.Log.Info("send photos to channel")
	for _, url := range urls {
		util.Log.Infof("sending: ", url)
		pic := tgbotapi.FileURL(url)
		msg := tgbotapi.NewPhotoToChannel(UserName, pic)
		_, err := tgbot.Bot.Send(msg)
		if err != nil {
			return err
		}
		util.Log.Infof("success,sleep 1 sec")
		time.Sleep(1 * time.Second)
	}
	util.Log.Infof("send photos to channel success")
	return nil
}
