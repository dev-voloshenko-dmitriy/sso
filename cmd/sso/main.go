package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"syscall"
)


const(
	envLocal = "local" 
	envDev = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("env", cfg))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.GRPCSrv.MustRun()

	// TODO: инициализация приложение (app)

	// TODO: запустить gRPC-сервер приложения

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCSrv.Stop()

	log.Info("application stoped")
}


func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {

	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)	
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)	
	}


	return log
}
