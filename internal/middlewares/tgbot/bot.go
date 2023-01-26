package tgbot

import (
	"time"

	"github.com/krau/favpics-helper/internal/structs"
	"github.com/krau/favpics-helper/pkg/util"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/krau/favpics-helper/pkg/config"
)

var TgBot *tgbotapi.BotAPI = new(tgbotapi.BotAPI)

func init() {
	util.Log.Info("init bot")
	if !config.Conf.Middlewares.TelegramBot.Enabled {
		util.Log.Info("bot disabled")
		return
	}
	bot, err := tgbotapi.NewBotAPI(config.Conf.Middlewares.TelegramBot.Token)
	util.Log.Infof("Login Tg Bot: [%s]", bot.Self.UserName)
	if err != nil {
		util.Log.Errorf("init bot error: %v", err)
	}
	TgBot = bot
	util.Log.Info("bot init success")
}

func SendPicsToChan(UserName string, pics []structs.Pic) error {
	util.Log.Info("send photos to tg channel")
	for _, pic := range pics {
		util.Log.Info("sending: ", pic.Title)
		mediaGroup := make([]interface{}, 0)
		if len(pic.Srcs) > 1 {
			firstPic := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(pic.Srcs[0]))
			firstPic.Caption = pic.Link
			mediaGroup = append(mediaGroup, firstPic)
			for _, src := range pic.Srcs[1:] {
				fileURL := tgbotapi.FileURL(src)
				tgPic := tgbotapi.NewInputMediaPhoto(fileURL)
				tgPic.Caption = pic.Link
				mediaGroup = append(mediaGroup, tgPic)
			}
			mediaConfig := tgbotapi.NewMediaGroup(0, mediaGroup)
			mediaConfig.ChannelUsername = UserName
			_, err := TgBot.SendMediaGroup(mediaConfig)
			if err != nil {
				util.Log.Errorf("send photo error: %v", err)
				return err
			}
			util.Log.Info("success,sleep 3 sec")
			time.Sleep(3 * time.Second)
		} else {
			tgPic := tgbotapi.FileURL(pic.Srcs[0])
			msg := tgbotapi.NewPhotoToChannel(UserName, tgPic)
			markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL(pic.Title, pic.Link)))
			msg.ReplyMarkup = markup
			_, err := TgBot.Send(msg)
			if err != nil {
				util.Log.Errorf("send photo error: %v", err)
				return err
			}
			util.Log.Infof("send photos to channel done")
		}
	}
	return nil
}
