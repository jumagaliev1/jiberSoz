package app

import (
	"context"
	"github.com/jumagaliev1/jiberSoz/internal/service"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	http "github.com/jumagaliev1/jiberSoz/internal/transport"
	"github.com/jumagaliev1/jiberSoz/internal/transport/http/handler"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	repo, err := storage.New(ctx)
	if err != nil {
		return err
	}

	svc, err := service.New(repo)
	if err != nil {
		return err
	}

	handl, err := handler.New(svc)
	if err != nil {
		return err
	}

	HTTPServer := http.NewServer(handl)

	return HTTPServer.StartHTTPServer(ctx)
}
