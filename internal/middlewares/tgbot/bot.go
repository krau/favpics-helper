package tgbot

import (
	"time"

	"github.com/krau/favpics-helper/internal/db"
	"github.com/krau/favpics-helper/internal/models"
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

/*
发送图片到频道，如果图片是多图，使用MediaGroup发送，否则使用Photo发送，
如果发送失败，使用Markdown发送图片链接。
如果仍然失败，返回错误。
*/
func SendPicsToChan(UserName string, pic models.Pic) error {
	util.Log.Info("sending pic to tg channel", pic.Link)
	if len(pic.Srcs) <= 1 {
		util.Log.Debug("a single photo: ", pic.Title)
		return sendPicToChan(UserName, pic)
	} else {
		util.Log.Debug("a pic group: ", pic.Title)
		mediaGroup := make([]interface{}, 0)
		firstPic := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(pic.Srcs[0]))
		firstPic.Caption = pic.Link + "\n" + pic.Title + "\n" + pic.Description
		mediaGroup = append(mediaGroup, firstPic)
		for _, src := range pic.Srcs[1:] {
			fileURL := tgbotapi.FileURL(src)
			tgPic := tgbotapi.NewInputMediaPhoto(fileURL)
			mediaGroup = append(mediaGroup, tgPic)
		}
		mediaConfig := tgbotapi.NewMediaGroup(0, mediaGroup)
		mediaConfig.ChannelUsername = UserName
		_, err := TgBot.SendMediaGroup(mediaConfig)
		if err != nil {
			util.Log.Noticef("send media group [%s] error: %v, try send single pic", pic.Link, err)
			return sendPicToChan(UserName, pic)
		}
		db.AddPic(pic)
		util.Log.Infof("%s sent,sleep 10 sec", pic.Title)
		time.Sleep(10 * time.Second)
	}
	return nil
}

func sendPicToChan(UserName string, pic models.Pic) error {
	util.Log.Debug("send photo to tg channel")
	msg := tgbotapi.NewPhotoToChannel(UserName, tgbotapi.FileURL(pic.Srcs[0]))
	markup := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL(pic.Title, pic.Link)))
	msg.ReplyMarkup = markup
	_, err := TgBot.Send(msg)
	if err != nil {
		util.Log.Noticef("send photo error: %v,try just send link", err)
		return sendPicLinkToChan(UserName, pic)
	}
	db.AddPic(pic)
	util.Log.Infof("%s sent,sleep 10 sec", pic.Title)
	time.Sleep(10 * time.Second)
	util.Log.Debug("send pic to channel done")
	return nil
}

func sendPicLinkToChan(UserName string, pic models.Pic) error {
	util.Log.Debug("send pic link to tg channel")
	picLink := "[" + pic.Title + "](" + pic.Link + ")"
	msg := tgbotapi.NewMessageToChannel(UserName, picLink)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = false
	_, err := TgBot.Send(msg)
	if err != nil {
		util.Log.Errorf("send pic link error: %v", err)
		return err
	}
	db.AddPic(pic)
	util.Log.Infof("%s sent,sleep 10 sec", pic.Title)
	time.Sleep(10 * time.Second)
	util.Log.Debug("send pic link to channel done")
	return nil
}
