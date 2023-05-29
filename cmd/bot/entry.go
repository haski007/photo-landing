package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/haski007/photo-landing/internal/app/bot/config"
	"github.com/haski007/photo-landing/internal/app/bot/service"
	"github.com/haski007/photo-landing/pkg/graceful"
	"github.com/haski007/photo-landing/pkg/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pb "github.com/haski007/photo-landing/api"
)

func Run(ctx context.Context, args run.Args) error {
	log := logrus.New()
	log.SetLevel(args.LogLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	var cfg config.Config
	if err := config.Load(args.ConfigFile, &cfg); err != nil {
		return fmt.Errorf("load config %s err: %w", args.ConfigFile, err)
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	// ---> Metrics
	httpMux := http.NewServeMux()
	httpMux.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{Addr: args.MetricsAddr, Handler: httpMux}

	// ---> Telegram
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramBot.Token)
	if err != nil {
		return fmt.Errorf("new tg bot api err: %w", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.TelegramBot.UpdatesTimeoutSec
	chUpdates := botApi.GetUpdatesChan(u)

	botSrv := service.NewBotPublisher(
		ctx,
		botApi,
		cfg.TelegramBot.CreatorUserID,
		chUpdates,
		cfg.WhiteList,
	).SetLogger(log)

	if err := tgbotapi.SetLogger(log); err != nil {
		return fmt.Errorf("set looger for tgbotapi package err: %w", err)
	}

	var services errgroup.Group

	services.Go(func() error {
		defer stop()

		me, er := botApi.GetMe()
		if er != nil {
			logrus.WithError(err).Println("bot api getMe")
		}

		log.Infof("bot @%s is polling now", me.UserName)

		if errL := botSrv.StartPool(); errL != nil {
			logrus.WithError(err).Println("bot listener exit with error")
		}

		return nil
	})

	grpcServer := grpc.NewServer()
	services.Go(func() error {
		defer stop()

		lis, err := net.Listen("tcp", cfg.GRPC.Port)
		if err != nil {
			return fmt.Errorf("grps setup failed to listen: %w", err)
		}

		pb.RegisterBotServiceServer(grpcServer, botSrv)

		reflection.Register(grpcServer)

		log.Infof("grpc service listening on %s", cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("grpc server serve err: %w", err)
		}

		return nil
	})

	services.Go(func() error {
		defer stop()
		log.Infof("metrics service listening on %s", args.MetricsAddr)

		if errLA := metricsServer.ListenAndServe(); errLA != nil && !errors.Is(errLA, http.ErrServerClosed) {
			logrus.WithError(err).Error("metrics services exit with error")
		}

		return nil
	})

	go graceful.Shutdown(
		ctx,
		//graceful.TGBOT(botSrv),
		graceful.GRPC(grpcServer),
		graceful.HTTP(metricsServer),
		graceful.CloseFunc(func() error {
			stop()
			log.Infof("services gracefully stopped")
			return nil
		}),
	)

	return services.Wait()
}
