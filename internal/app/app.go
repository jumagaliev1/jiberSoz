package app

import (
	"context"
	pb "github.com/jumagaliev1/jiberSoz/hasher/proto"
	"github.com/jumagaliev1/jiberSoz/internal/service"
	"github.com/jumagaliev1/jiberSoz/internal/storage"
	"github.com/jumagaliev1/jiberSoz/internal/storage/redis"
	http "github.com/jumagaliev1/jiberSoz/internal/transport"
	"github.com/jumagaliev1/jiberSoz/internal/transport/http/handler"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.Dial(viper.GetString("grpc.uri"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	grpcServer := pb.NewHashServiceClient(conn)

	cacheView := redis.New(viper.GetString("redis.view.uri"))
	cachePost := redis.New(viper.GetString("redis.post.uri"))

	svc, err := service.New(repo, cacheView, cachePost, grpcServer)
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
