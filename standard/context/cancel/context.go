package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			break
		}
		fmt.Println("g1 done")
	}(ctx)

	ctx2, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			break
		}
		fmt.Println("g2 done")
	}(ctx2)

	time.Sleep(time.Second * 5)
	cancel()

	// init signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sc
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
