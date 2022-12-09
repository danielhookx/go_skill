package benchmark

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"sync/atomic"

	"github.com/oofpgDLD/go_skill/exercise/net/proto"
)

type Client struct {
	ch       chan string
	conn     net.Conn
	isClosed int32
}

func NewClient() *Client {
	c := &Client{
		ch:       make(chan string, 100),
		isClosed: 0,
	}
	return c
}

func (c *Client) Connect() (err error) {
	// Simply check that the server is up and can
	// accept connections.
	c.conn, err = net.Dial("tcp", "localhost:11123")
	if err != nil {
		return
	}
	go c.serve()
	return
}

func (c *Client) serve() {
	//init
	var (
		rr  *bufio.Reader
		rw  *bufio.Writer
		err error
	)

	rr = bufio.NewReader(c.conn)
	rw = bufio.NewWriter(c.conn)
	//write
	go c.dispatchWebsocket(rw)

	//reader
	for true {
		p := proto.Frame{}
		err = proto.Read(rr, &p)
		if err != nil {
			break
		}
		log.Println(string(p.Body))
	}
	//release
	if err != nil && err != io.EOF && !strings.Contains(err.Error(), "closed") {
		log.Println("server tcp failed:", err)
	}
	c.Close()
}

func (c *Client) dispatchWebsocket(rw *bufio.Writer) {
	var (
		err error
	)
	for {
		select {
		case s := <-c.ch:
			p := proto.Frame{}
			p.Body = []byte(s)
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
	c.Close()
}

func (c *Client) Push(s string) {
	c.ch <- s
}

func (c *Client) Close() error {
	if atomic.CompareAndSwapInt32(&c.isClosed, 0, 1) {
		return c.conn.Close()
	}
	return nil
}
