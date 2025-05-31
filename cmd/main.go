package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/critiq/crypto_bot/api"
	"github.com/critiq/crypto_bot/buttons"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// logging
		user := update.Message.From
		userName := user.UserName
		firstName := user.FirstName
		lastName := user.LastName
		text := update.Message.Text

		log.Printf("@%s (%s %s): %s", userName, firstName, lastName, text)

		if update.Message.Text == "/start" {
			name := update.Message.From.FirstName
			msgText := "Hello, " + name + "\n" +
				"I'm crypto price bot. To check the price of a coin, just send:\n" +
				"`/price BTC`\n" + "or other coin"

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			msg.ReplyMarkup = buttons.MainMenu()
			msg.ParseMode = "Markdown"
			bot.Send(msg)
			continue
		}

		if strings.HasPrefix(update.Message.Text, "/convert") {
			parts := strings.Split(update.Message.Text, " ")
			if len(parts) != 4 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Usage: /convert [amout] [from] [to]"))
				continue
			}

			amount, _ := strconv.ParseFloat(parts[1], 64)
			from := strings.ToLower(parts[2])
			to := strings.ToLower(parts[3])

			converted, err := api.GetConvert(from, to, amount)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Conversion error: "+err.Error()))
				continue
			}

			response := fmt.Sprintf("%.4f %s = %.4f %s", amount, from, converted, to)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
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

			log.Printf(bot.Self.FirstName, reply)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}
	}
}
