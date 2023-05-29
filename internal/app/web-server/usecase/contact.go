package usecase

import (
	"context"
	"fmt"
	pb "github.com/haski007/photo-landing/api"
)

type ContactUseCase struct {
	botClient pb.BotServiceClient
}

func NewContactUseCase(botClient pb.BotServiceClient) *ContactUseCase {
	return &ContactUseCase{
		botClient: botClient,
	}
}

func (c *ContactUseCase) SendContactForm(name, email, message string) error {
	rsp, err := c.botClient.SendContactRequest(context.Background(), &pb.ContactRequest{
		Name:    name,
		Email:   email,
		Message: message,
	})
	if err != nil {
		return fmt.Errorf("SendContactRequest err: %w", err)
	}

	if !rsp.GetOk() {
		return fmt.Errorf("SendContactRequest errors: %s", rsp.GetErrors())
	}
	return nil
}
