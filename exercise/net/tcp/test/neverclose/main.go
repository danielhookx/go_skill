package main

import (
	"fmt"
	"log"
	"net"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:11123")
	if err != nil {
		panic(err)
	}

	count := 1000
	cons := make([]net.Conn, count)
	for i := 0; i < count; i++ {
		conn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			fmt.Printf("net.DialTCP %d failed err=%s", i, err.Error())
		}
		cons[i] = conn
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
