package handlers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	t.Run("тест ответа 200", func(t *testing.T) {
		w := httptest.NewRecorder()

		SendResponse(w, struct {
			Test string `json:"test"`
		}{Test: "test"})

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{\"test\":\"test\"}\n", w.Body.String())
	})

	t.Run("тест ответа 400", func(t *testing.T) {
		w := httptest.NewRecorder()

		SendError(w, errors.New("test error"))

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, "{\"error\":\"test error\"}\n", w.Body.String())
	})
}
