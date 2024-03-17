package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func doSomeJobs() {
	fmt.Println("do some jobs")
}

func main() {
	sigalChan := make(chan os.Signal, 1)
	signal.Notify(sigalChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigalChan
	doSomeJobs()
}
