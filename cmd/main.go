package main

import (
	"log"
	"strings"

	"github.com/critiq/crypto_bot/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("7563721679:AAEXPwmSC2lbg4E5h4V1_LA86OSWt5Jearo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", &bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			name := update.Message.From.FirstName
			msgText := "Hello, " + name + "\n\n" +
				"I'm crypto price bot. To check the price of a coin, just send:\n" +
				"`/price BTC`\n" + "or other coin"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			msg.ParseMode = "Markdown"
			bot.Send(msg)
			continue
		}
		if strings.HasPrefix(update.Message.Text, "/price") {
			args := strings.Split(update.Message.Text, " ")

			if len(args) < 2 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Write symbol coin, /price symbol")
				bot.Send(msg)
				continue
			}
			symbol := strings.ToUpper(args[1]) + "USDT"

			price, err := api.GetPrice(symbol)
			var reply string

			if err != nil {
				reply = "Symbol not found"
			} else {
				reply = "Price " + symbol + ": $" + price
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}

	}
}
