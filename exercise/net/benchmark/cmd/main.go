package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oofpgDLD/go_skill/exercise/net/benchmark"
)

func main() {
	userNum := 200
	for i := 0; i < userNum; i++ {
		go func(uid int) {
			cli := benchmark.NewClient()
			err := cli.Connect()
			if err != nil {
				panic(err)
			}
			current := 0
			for true {
				current++
				cli.Push(fmt.Sprintf("user %d send hello %d", uid, current))
				time.Sleep(time.Millisecond * 1000)
			}
		}(i)
	}
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Println("service get a signal")
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
