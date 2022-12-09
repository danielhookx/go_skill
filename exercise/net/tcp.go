package net

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"strings"

	"github.com/oofpgDLD/go_skill/exercise/net/proto"
)

type TCPServer struct {
	l net.Listener
}

func NewTCPServer() *TCPServer {
	return &TCPServer{}
}

func (t *TCPServer) Run(addr string) error {
	var tcpAddr *net.TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	t.l, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	for {
		conn, err := t.l.(*net.TCPListener).AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		if conn == nil {
			log.Println(errors.New("conn is nil"))
			continue
		}
		// serve
		err = conn.SetKeepAlive(true)
		if err != nil {
			log.Println(err)
			continue
		}
		go t.handleTCPConn(conn)
	}
}

func (t *TCPServer) handleTCPConn(conn net.Conn) {
	//init
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
	//write
	go dispatchWebsocket(conn, rw, ch)

	//reader
	for true {
		p := proto.Frame{}
		err = proto.Read(rr, &p)
		if err != nil {
			break
		}
		select {
		case ch <- string(p.Body):
		default:
		}
		log.Println("server rev:", string(p.Body))
	}
	//release
	if err != nil && err != io.EOF && !strings.Contains(err.Error(), "closed") {
		log.Println("server tcp failed:", err)
	}
	conn.Close()
}

func dispatchWebsocket(conn net.Conn, rw *bufio.Writer, ch chan string) {
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
}

func (t *TCPServer) Stop() error {
	return t.l.Close()
}
