package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	xnet "github.com/danielhookx/go_skill/exercise/net/tcp"
)

func main() {
	// Start the new server.
	srv, err := xnet.ListenTCP(":11123", xnet.Reject)
	if err != nil {
		log.Println("error starting TCP server")
		return
	}
	log.Println("start serve :11123")
	srv.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Println("service get a signal")
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			srv.Stop()
			return
		case syscall.SIGHUP:
			// TODO reload
		default:
			return
		}
	}
}
