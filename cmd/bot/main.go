package main

import (
	"github.com/Muravjeva/route_256_t_me_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	commander := NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			case "withArgument":
				msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
			case "html":
				msg.ParseMode = "html"
				msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
			case "list":
				commander.List(update.Message)
			default:
				commander.Default(update.Message)
			}

			bot.Send(msg)
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)
		bot.Send(msg)
	}
}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI,
	productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) Help(inputMssage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMssage.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	c.bot.Send(msg)
}

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

func (c *Commander) Default(inputMssage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMssage.From.UserName, inputMssage.Text)
	msg := tgbotapi.NewMessage(inputMssage.Chat.ID, "You wrote: "+inputMssage.Text)
	c.bot.Send(msg)
}
