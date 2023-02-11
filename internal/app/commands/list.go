package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMssage *tgbotapi.Message) {
	outputMsgText := "All products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMssage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}
