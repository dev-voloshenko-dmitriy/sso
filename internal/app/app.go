package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)


type App struct {
	GRPCSrv *grpcapp.App
}


func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	totenTTL time.Duration,
) *App {
	// TODO: инициализировать хранилище (storage)

	// TODO: init auth servese (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}


