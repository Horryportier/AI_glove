package server

import (
	"fmt"
	"io"
	"net"
)

type Message struct {
	From    string
	Payload []byte
}

type Server struct {
	ListenAddr string
	Ln         net.Listener
	Quitch     chan struct{}
	Msgch      chan Message
	DeviceIp   string
}

func NewServer(ListenAddr string) *Server {
	return &Server{
		ListenAddr: ListenAddr,
		Quitch:     make(chan struct{}),
		Msgch:      make(chan Message, 100),
	}
}

func (s *Server) Start() error {
	Ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	defer Ln.Close()
	s.Ln = Ln

	go s.acceptLoop()

	<-s.Quitch
	close(s.Msgch)
	close(s.Quitch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", ErrorStyle.Render(err.Error()))
			continue
		}
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("IO EOF:", ErrorStyle.Render(err.Error()))
			break
		}
		if err != nil {
			fmt.Println("read error:", ErrorStyle.Render(err.Error()))
			continue
		}

		s.Msgch <- Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf,
		}
		conn.Write([]byte("200"))
	}
}
