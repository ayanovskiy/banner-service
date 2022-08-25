package server

import (
	banners "banner-service/internal/server/handlers/banner"
	slots "banner-service/internal/server/handlers/slot"
	"banner-service/internal/storage"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
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

func (s *Server) Run(storage storage.IStorage) error {
	server := &http.Server{
		Addr:    s.addr,
		Handler: s.createRouters(storage),
	}

	fmt.Println("Server is starting...")
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) createRouters(storage storage.IStorage) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := io.WriteString(w, "main page")
		if err != nil {
			log.Fatal(err)
		}
	})
	router.HandleFunc("/slot/add-banner", func(w http.ResponseWriter, r *http.Request) {
		slots.AddBannerToSlot(storage, w, r)
	})
	router.HandleFunc("/slot/remove-banner", func(w http.ResponseWriter, r *http.Request) {
		slots.RemoveBannerFromSlot(storage, w, r)
	})
	router.HandleFunc("/banner/select", func(w http.ResponseWriter, r *http.Request) {
		banners.SelectBannerHandler(storage, w, r)
	})
	router.HandleFunc("/banner/hit", func(w http.ResponseWriter, r *http.Request) {
		banners.HitBannerRequest(storage, w, r)
	})

	return router
}
