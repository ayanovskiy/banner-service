package app

import (
	"banner-service/internal/rmq"
	"banner-service/internal/server"
	"banner-service/internal/storage"
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

type App struct {
	config  *Config
	server  *server.Server
	storage storage.IStorage
	rmq     *rmq.Rabbit
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
	rabbit, err := rmq.NewRabbit(ctx, cfg.Rabbit.Dsn, cfg.Rabbit.Exchange, cfg.Rabbit.Queue, cfg.Rabbit.Tag)
	if err != nil {
		return nil, err
	}

	return &App{
		config:  cfg,
		server:  server.NewServer(cfg.Server.Addr, cfg.Server.Port),
		storage: storage.NewStorage(ctx, db),
		rmq:     rabbit,
	}, nil
}

func (app App) Run() {
	defer app.storage.CloseDb()

	err := app.server.Run(server.NewRouters(app.storage, app.rmq))
	if err != nil {
		log.Fatal(err)
	}
}
