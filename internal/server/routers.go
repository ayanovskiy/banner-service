package server

import (
	"banner-service/internal/rmq"
	banners "banner-service/internal/server/handlers/banner"
	slots "banner-service/internal/server/handlers/slot"
	stats "banner-service/internal/server/handlers/stat"
	"banner-service/internal/storage"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func NewRouters(storage storage.IStorage, rabbit *rmq.Rabbit) *mux.Router {
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
	router.HandleFunc("/stats/send", func(w http.ResponseWriter, r *http.Request) {
		stats.SendStatHandler(storage, rabbit, w, r)
	})

	return router
}
