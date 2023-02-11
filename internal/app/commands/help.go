package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMssage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMssage.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	c.bot.Send(msg)
}
