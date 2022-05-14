package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)
	go func() {
		for update := range updates {
			go handleUpdate(bot, update)
		}
	}()
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := update.Message.Text
	chatID := update.Message.Chat.ID
	replyMsg := tgbotapi.NewMessage(chatID, text)
	replyMsg.ReplyToMessageID = update.Message.MessageID
	if update.Message.IsCommand() {

	} else {

	}

}
