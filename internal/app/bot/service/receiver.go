package service

import (
	"fmt"

	"github.com/haski007/photo-landing/pkg/emoji"
	"github.com/sirupsen/logrus"
)

func (rcv *BotPublisher) StartPool() error {
	me, err := rcv.bot.GetMe()
	if err != nil {
		_ = rcv.NotifyCreator(fmt.Sprintf("[bot GetMe] err: %s", err))
		return err
	}

	for update := range rcv.updates {
		if update.EditedMessage != nil || update.Poll != nil {
			continue
		}

		if update.PollAnswer != nil {
			continue
		}

		// Check if someone added bot to chat
		if update.MyChatMember != nil &&
			update.MyChatMember.NewChatMember.User.ID == me.ID {
			go func() {
				if err := rcv.sendStartInfo(update.MyChatMember.Chat.ID); err != nil {
					rcv.log.WithError(err).Println("[new chat member update] send start info")
				}
			}()
			continue
		}

		// ---> Commands
		if update.Message != nil && update.Message.IsCommand() {
			rcv.log.WithFields(map[string]interface{}{
				"from":    update.Message.From.UserName,
				"from_id": update.Message.From.ID,
				"message": update.Message.Text,
			}).Infoln("got command")

			command := update.Message.Command()
			switch {
			case command == "help":
				go rcv.cmdStartHandler(update)
			case command == "start":
				go rcv.cmdStartHandler(update)

			default:
				go func() {
					if err := rcv.SendMessage(
						update.Message.Chat.ID,
						"Such command does not exist! "+emoji.NoEntry,
					); err != nil {
						logrus.WithError(err).Printf("send message to chat: %d", update.Message.Chat.ID)
					}
				}()
			}
		}

		// Parse messages
		if update.Message != nil && !update.Message.IsCommand() {
			rcv.log.WithFields(map[string]interface{}{
				"from":    update.Message.From.UserName,
				"from_id": update.Message.From.ID,
				"message": update.Message.Text,
			}).Infoln("got message")
			switch {
			}
		}
	}

	logrus.Printf("Channel is closed")
	return nil
}
