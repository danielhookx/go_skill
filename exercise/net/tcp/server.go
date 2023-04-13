package tcp

import (
	"errors"
	"log"
	"net"
	"runtime/debug"
)

type Server struct {
	l        net.Listener
	acceptCB func(conn net.Conn)
}

func ListenTCP(addr string, cb func(conn net.Conn)) (*Server, error) {
	var laddr *net.TCPAddr
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	l, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return nil, err
	}
	return &Server{
		l:        l,
		acceptCB: cb,
	}, nil
}

func (s *Server) Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("tcp listener accept failed:%+v\n\n%s", r, debug.Stack())
			}
		}()

		for {
			conn, err := s.l.(*net.TCPListener).AcceptTCP()
			if err != nil {
				log.Println(err)
				continue
			}
			if conn == nil {
				log.Println(errors.New("conn is nil"))
				continue
			}
			go s.acceptCB(conn)
		}
	}()
}

func (s *Server) Stop() error {
	return s.l.Close()
}
