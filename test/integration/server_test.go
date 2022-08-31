package integration

import (
	"banner-service/internal/server"
	"context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("тест работы сервера", func(t *testing.T) {
		r := mux.NewRouter()
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			_, err := io.WriteString(w, "main page")
			if err != nil {
				t.Error(err)
			}
		})

		ctx := context.Background()

		srv := server.NewServer("localhost", 1111)

		go func() {
			err := srv.Run(r)
			if err != nil {
				t.Error(err)
			}
		}()

		c := &http.Client{
			Timeout: 60 * time.Second,
		}
		request, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:1111/", nil)
		resp, err := c.Do(request)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "main page", string(body))
	})
}
