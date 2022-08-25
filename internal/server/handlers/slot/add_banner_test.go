package handlers

import (
	"banner-service/internal/model"
	"banner-service/test/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddBannerToSlot(t *testing.T) {
	t.Run("корректное добавление баннера", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		s := mock.NewMockIStorage(ctrl)
		gomock.InOrder(
			s.EXPECT().FindSlot(uint(1)).Return(&model.Slot{Id: 1, Desc: "Test 1"}, nil),
			s.EXPECT().FindBanner(uint(1)).Return(&model.Banner{Id: 1, Desc: "Test 1"}, nil),
			s.EXPECT().AddBannerToSlot(uint(1), uint(1)).Return(nil),
		)

		req, err := http.NewRequest(
			http.MethodPost,
			"/slot/add-banner",
			strings.NewReader(`{"slot_id":1,"banner_id":1}`),
		)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		AddBannerToSlot(s, rr, req)

		expectedResult := "{\"response\":\"ok\"}\n"
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedResult, rr.Body.String())
	})

	t.Run("перехват ошибки при добавлении баннера", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		s := mock.NewMockIStorage(ctrl)
		gomock.InOrder(
			s.EXPECT().FindSlot(uint(15)).Return(nil, errors.New("no rows")),
		)

		req, err := http.NewRequest(
			http.MethodPost,
			"/slot/add-banner",
			strings.NewReader(`{"slot_id":15,"banner_id":1}`),
		)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		AddBannerToSlot(s, rr, req)

		expectedResult := "{\"error\":\"no rows\"}\n"
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expectedResult, rr.Body.String())
	})
}
