package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) Default(inputMssage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMssage.From.UserName, inputMssage.Text)
	msg := tgbotapi.NewMessage(inputMssage.Chat.ID, "You wrote: "+inputMssage.Text)
	c.bot.Send(msg)
}
