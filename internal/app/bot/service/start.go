package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/haski007/photo-landing/pkg/emoji"
)

func (rcv *BotPublisher) cmdStartHandler(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := rcv.sendStartInfo(chatID); err != nil {
		rcv.log.WithError(err).Println("[cmdStartHandler] send start info")
	}
}

func (rcv *BotPublisher) sendStartInfo(chatID int64) error {
	message := `Hello, I am here to publish users contact requests for you. ` + emoji.FaceWinking
	return rcv.SendMessage(chatID, message)
}
