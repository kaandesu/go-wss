package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	addr  string
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWs(wsConn *websocket.Conn) {
	fmt.Println("Incoming request from:", wsConn.RemoteAddr().String())
	s.conns[wsConn] = true

	go s.handleConnection(wsConn)
}

func (s *Server) handleConnection(wsConn *websocket.Conn) {
	buf := make([]byte, 1024)
	defer wsConn.Close()
	for {
		n, err := wsConn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read error: ", err)
			break
		}

		payload := buf[:n]
		fmt.Println(string(payload))
		wsConn.Write([]byte("msg recieved!\n"))
	}
}

func main() {
	server := NewServer("127.0.0.1:3000")
	http.Handle("/wss", websocket.Handler(server.handleWs))
	http.ListenAndServeTLS(":3000", "./server.crt", "./server.key", nil)
}
