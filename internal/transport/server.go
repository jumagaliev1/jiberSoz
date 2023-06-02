package transport

import (
	"context"
	"fmt"
	"github.com/jumagaliev1/jiberSoz/internal/transport/http/handler"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	App     *echo.Echo
	Handler *handler.Handler
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{Handler: handler}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()
	go func() {
		if err := s.App.Start(fmt.Sprintf(":%v", viper.GetString("port"))); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:#{err}\n")
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:#{err}")
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Use(echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
			log.Info(map[string]interface{}{"URI": v.URI, "status": v.Status})
			return nil
		},
	}))

	return e
}
