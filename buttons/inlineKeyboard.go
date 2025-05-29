package buttons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func MainMenu() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("BTC", "price_btc"),
			tgbotapi.NewInlineKeyboardButtonData("ETH", "price_eth"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("Convert", "convert"),
		},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
