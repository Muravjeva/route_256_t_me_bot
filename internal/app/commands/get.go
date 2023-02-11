package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMssage *tgbotapi.Message) {
	args := inputMssage.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("Fail to get product with idx: %d: %v ", idx, err)
		return
	}
	msg := tgbotapi.NewMessage(inputMssage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
