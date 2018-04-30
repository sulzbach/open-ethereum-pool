package alerts

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

type TelegramConfig struct {
	Enabled bool   `json:"enabled"`
	Api     string `json:"api"`
	Chat    int64  `json:"chat"`
}

type TelegramBot struct {
	config   *TelegramConfig
	bot      *tgbotapi.BotAPI
	halt     bool
	lastFail error
}

func NewTelegramBot(cfg *TelegramConfig) *TelegramBot {
	u := &TelegramBot{config: cfg}
	u.bot, u.lastFail = tgbotapi.NewBotAPI(cfg.Api)
	if u.lastFail != nil {
		log.Panic(u.lastFail)
	}
	u.bot.Debug = true
	return u
}

func (u *TelegramBot) SendMessage(text string) bool {
	msg := tgbotapi.NewMessage(u.config.Chat, text)
	//	msg.ReplyToMessageID = 11
	u.bot.Send(msg)
	return true
}
