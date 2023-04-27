package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yamk12nfu/toy-zaim/app/drivers/infrastructure"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	router := infrastructure.NewRouter()
	router.Start()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	router.Shutdown(ctx)
}
