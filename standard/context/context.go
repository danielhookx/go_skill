package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		wg.Add(1)
		do(ctx)
		wg.Done()
	}()

	// init signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sc
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			wg.Wait()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func do(ctx context.Context) {
	num := 60

	for i := 0; i < num; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}
