package main

import (
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

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)
			//switch update.Message.Command() {
			//case "help":
			//	msg.Text = "type /sayhi or /status."
			//case "sayhi":
			//	msg.Text = "Hi :)"
			//case "status":
			//	msg.Text = "I'm ok."
			//case "withArgument":
			//	msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
			//case "html":
			//	msg.ParseMode = "html"
			//	msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
			//default:
			//	msg.Text = "I don't know that command"
			//}
			bot.Send(msg)
		}

	}
}
