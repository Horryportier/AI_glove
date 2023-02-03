package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

        s := NewServer(":8090")

	go func() {
		for msg := range s.msgch {
			fmt.Printf("msg from conn (%v): %s\n", 
                        msg.from, string(msg.payload))
		}
	}()

        log.Fatal(s.Start())
}

type Message struct {
        from string 
        payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message ),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
        close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		fmt.Println("new connection to the server:", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
        buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		s.msgch <- Message{
                        from: conn.RemoteAddr().String(),
                        payload: buf[:n],
                }
	}
}
