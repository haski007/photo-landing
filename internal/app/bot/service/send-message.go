package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (rcv *BotPublisher) SendMessage(chatID int64, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = tgbotapi.ModeMarkdown

	_, err := rcv.bot.Send(message)
	return err
}

func (rcv *BotPublisher) Reply(chatID int64, messageID int, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = tgbotapi.ModeMarkdown
	message.ReplyToMessageID = messageID

	_, err := rcv.bot.Send(message)
	return err
}

func (rcv *BotPublisher) SendMessageWithoutMarkdown(chatID int64, text string) error {
	message := tgbotapi.NewMessage(chatID, text)

	_, err := rcv.bot.Send(message)
	return err
}
