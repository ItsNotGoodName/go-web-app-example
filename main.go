package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/ItsNotGoodName/go-web-app-example/web"
	"github.com/ItsNotGoodName/go-web-app-example/web/events"
	"github.com/ItsNotGoodName/go-web-app-example/web/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Address         string
	ShutdownTimeout time.Duration
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg := Config{
		Address:         ":8080",
		ShutdownTimeout: 5 * time.Second,
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: web.AssetFS(),
	}))
	e.Use(web.MetaMiddleware())

	// Handlers
	e.GET("/", func(c echo.Context) error {
		return web.Render(c, pages.Index("Hello World!", cfg))
	})
	e.POST("/", func(c echo.Context) error {
		events.Toast("Toast from server!").SetTrigger(c.Response())
		return nil
	})

	start(ctx, e, cfg.Address, cfg.ShutdownTimeout)
}

func start(ctx context.Context, e *echo.Echo, address string, shutdownTimeout time.Duration) {
	errC := make(chan error, 1)
	go func() { errC <- e.Start(address) }()

	web.ReloadVite()

	select {
	case err := <-errC:
		log.Fatalln(err)
	case <-ctx.Done():
		slog.Info("Gracefully shutting down HTTP server...")

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			slog.Error("HTTP Server failed to shutdown gracefully", "error", err)
		}
	}
}
