package server

import (
	"fmt"
	"net"

)
type Message struct {
        From string 
        Payload []byte
}

type Server struct {
	ListenAddr string
	Ln         net.Listener
	Quitch     chan struct{}
	Msgch      chan Message
}

func NewServer(ListenAddr string) *Server {
	return &Server{
		ListenAddr: ListenAddr,
		Quitch:     make(chan struct{}),
		Msgch:      make(chan Message, 100),
	}
}

func (s *Server) Start() error {
	Ln, err := net.Listen("tcp", "0.0.0.0"+s.ListenAddr)
	if err != nil {
		return err
	}
	defer Ln.Close()
	s.Ln = Ln

	go s.acceptLoop()

	<-s.Quitch
        close(s.Msgch)
	

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}

		//fmt.Println("new connection to the server:", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
        buf := make([]byte, 1024)
	for {
		conn.Read(buf)
                //fmt.Println("read error:", string(buf))
		//if err != nil {
		//	fmt.Println("read error:", err)
		//	//continue
		//}

		s.Msgch <- Message{
                        From: conn.RemoteAddr().String(),
                        Payload: buf,
                }

                conn.Write([]byte("arigatougozaimasu"))
	}
}

