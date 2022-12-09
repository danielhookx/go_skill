package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	xnet "github.com/oofpgDLD/go_skill/exercise/net"
)

func main() {
	// Start the new server.
	srv := xnet.NewTCPServer()

	go func() {
		err := srv.Run(":11123")
		if err != nil {
			log.Println("error starting TCP server")
			return
		}
	}()
	log.Println("start serve :11123")

	if err := http.ListenAndServe(":8801", http.DefaultServeMux); err != nil {
		panic(err)
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
