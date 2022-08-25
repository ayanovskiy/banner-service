package app

import (
	"banner-service/internal/server"
	"banner-service/internal/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type App struct {
	config  *Config
	ctx     context.Context
	server  *server.Server
	storage storage.IStorage
}

func NewApp(configPath string) (*App, error) {
	cfg, err := NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	db, err := pgx.Connect(ctx, cfg.Database.Dsn)
	if err != nil {
		return nil, err
	}

	return &App{
		config:  cfg,
		server:  server.NewServer(cfg.Server.Addr, cfg.Server.Port),
		storage: storage.NewStorage(ctx, db),
	}, nil
}

func (app App) Run() {
	go func() {
		<-app.ctx.Done()
		defer app.storage.CloseDb()
	}()

	err := app.server.Run(app.storage)
	if err != nil {
		fmt.Println(err)
	}
}
