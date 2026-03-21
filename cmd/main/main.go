package main

import (
	"bloger/cmd/bootstrap"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	app, err := bootstrap.Init()
	if err != nil {
		log.Fatal(err)
		return
	}

	go func() {
		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Down(ctx); err != nil {
		log.Printf("应用关闭失败: %v", err)
	}
}
