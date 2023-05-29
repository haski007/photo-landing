package service

import (
	"context"
	"fmt"
	pb "github.com/haski007/photo-landing/api"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotPublisher struct {
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel

	creatorID int64
	log       logrus.FieldLogger
	receivers []int64

	ctx context.Context
	pb.UnimplementedBotServiceServer
}

func NewBotPublisher(
	ctx context.Context,
	botApi *tgbotapi.BotAPI,
	creatorID int64,
	updatesChan tgbotapi.UpdatesChannel,
	receivers []int64,
) *BotPublisher {
	return &BotPublisher{
		bot:       botApi,
		creatorID: creatorID,
		updates:   updatesChan,
		ctx:       ctx,
		receivers: receivers,
	}
}

func (rcv *BotPublisher) SetLogger(logger logrus.FieldLogger) *BotPublisher {
	rcv.log = logger.WithField("handler", "rcv")
	return rcv
}

func (rcv *BotPublisher) NotifyCreator(message string) error {
	if err := rcv.SendMessage(rcv.creatorID, message); err != nil {
		return fmt.Errorf("notify creator err: %w", err)
	}
	return nil
}
