package service

import (
	"context"
	"fmt"

	"github.com/haski007/photo-landing/pkg/emoji"

	pb "github.com/haski007/photo-landing/api"
)

func (rcv *BotPublisher) SendContactRequest(ctx context.Context, req *pb.ContactRequest) (*pb.ContactResponse, error) {
	msg := "Got new contact request " + emoji.Envelope + "\n"
	msg += fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", req.GetName(), req.GetEmail(), req.GetMessage())

	var errors []string
	for _, receiver := range rcv.receivers {
		if err := rcv.SendMessageWithoutMarkdown(receiver, msg); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) != 0 {
		return &pb.ContactResponse{Ok: false, Errors: errors}, nil
	}

	return &pb.ContactResponse{Ok: true}, nil
}
