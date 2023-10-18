package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	// CONNECTION TO REDIS
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis. \n Msg: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("[ERROR] Failed to Close Redis.\n Msg:", err)
		}
	}()

	fmt.Println("[INFO] Starting Web Server ... ")

	// GO CHANNEL
	// TODO: RESEARCH THIS
	ch := make(chan error, 1)

	// START WEB SERVER CONCURRENTLY
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start the web server. \n Msg: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		// GRACEFULLY SHUTDOWN SERVER AND TIMEOUT IN 10 SECONDS (ALLOW REQUESTS TO RESOLVE)
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
