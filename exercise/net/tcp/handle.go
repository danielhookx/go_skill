package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"

	"github.com/oofpgDLD/go_skill/exercise/net/tcp/proto"
)

func Reject(conn net.Conn) {
	log.Printf("reject pre conn and close %s", conn.RemoteAddr().String())
	conn.Close()
}

func EchoString(conn net.Conn) {
	var (
		rr  *bufio.Reader
		rw  *bufio.Writer
		err error
		ch  chan string
	)

	ch = make(chan string, 100)
	rr = bufio.NewReader(conn)
	rw = bufio.NewWriter(conn)
	log.Println("accepted:", conn.RemoteAddr())

	//write goroutine
	go func(conn net.Conn, rw *bufio.Writer, ch chan string) {
		var (
			err error
		)
		for {
			select {
			case s := <-ch:
				p := proto.Frame{}
				p.Body = []byte("server echo:" + s)
				err = proto.Write(rw, &p)
				if err != nil {
					goto failed
				}
				err = rw.Flush()
				if err != nil {
					goto failed
				}
			}
		}
	failed:
		if err != nil {
			log.Println("dispatch tcp failed:", err)
		}
		conn.Close()
	}(conn, rw, ch)

	//read goroutine
	for true {
		p := proto.Frame{}
		err = proto.Read(rr, &p)
		if err != nil {
			break
		}
		// blocked while local channel full
		select {
		case ch <- string(p.Body):
		}
		log.Println("server rev:", string(p.Body))
	}
	//release
	if err != nil && err != io.EOF && !strings.Contains(err.Error(), "closed") {
		log.Println("server tcp failed:", err)
	}
	conn.Close()
}
