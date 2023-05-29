package main

import (
	pb "github.com/haski007/photo-landing/api"
	"github.com/haski007/photo-landing/internal/app/web-server/config"
	"github.com/haski007/photo-landing/internal/app/web-server/handler"
	"github.com/haski007/photo-landing/internal/app/web-server/server"
	"github.com/haski007/photo-landing/internal/app/web-server/usecase"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	cc, err := grpc.Dial(cfg.BotServiceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("failed to dial bot service: %s", err.Error())
	}

	contactUseCase := usecase.NewContactUseCase(pb.NewBotServiceClient(cc))
	h := handler.NewHandler(contactUseCase)

	s := server.NewServer(h, cfg.Port)

	if err := s.Run(); err != nil {
		logrus.Fatalf("error running server: %s", err.Error())
	}
}
