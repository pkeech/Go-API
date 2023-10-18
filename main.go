package main

import (
	"context"
	"fmt"

	application "github.com/pkeech/Go-API/app"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("[ERROR] Failed to Start Application. \n Error:", err)
	}
}
