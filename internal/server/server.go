package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string, port uint) *Server {
	return &Server{
		addr: fmt.Sprintf("%s:%d", addr, port),
	}
}

func (s *Server) Run(routers *mux.Router) error {
	server := &http.Server{
		Addr:    s.addr,
		Handler: routers,
	}

	fmt.Println("Server is starting...")
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
