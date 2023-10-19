package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	application "github.com/pkeech/Go-API/app"
)

func main() {
	app := application.New(application.LoadConfig())

	// CREATE CONTEXT TO HANDLE GRACEFUL SHUTDOWNS
	// OS.INTERRUPT == SIGINIT
	// DEFER CANCEL MEANS IT WILL RUN AT THE END
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("[ERROR] Failed to Start Application. \n Error:", err)
	}

}
