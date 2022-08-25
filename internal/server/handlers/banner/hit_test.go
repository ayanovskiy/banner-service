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

func TestHitBannerRequest(t *testing.T) {
	t.Run("корректный переход по баннеру", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		s := mock.NewMockIStorage(ctrl)
		gomock.InOrder(
			s.EXPECT().FindSlot(uint(1)).Return(&model.Slot{Id: 1, Desc: "Test 1"}, nil),
			s.EXPECT().FindBanner(uint(1)).Return(&model.Banner{Id: 1, Desc: "Test 1"}, nil),
			s.EXPECT().FindGroup(uint(1)).Return(&model.Group{Id: 1, Desc: "Test 1"}, nil),
			s.EXPECT().FindOrCreateStat(uint(1), uint(1), uint(1)).Return(&model.Stat{Id: 1, SlotId: 1, BannerId: 1, GroupId: 1, Shows: 1}, nil),
			s.EXPECT().UpdateStat(uint(1), uint(1), uint(1)),
		)

		req, err := http.NewRequest(
			http.MethodPost,
			"/banner/hit",
			strings.NewReader(`{"slot_id":1,"banner_id":1,"group_id":1}`),
		)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		HitBannerRequest(s, rr, req)

		expectedResult := "{\"response\":\"ok\"}\n"
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedResult, rr.Body.String())
	})

	t.Run("перехват ошибки при переходе по баннеру", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		s := mock.NewMockIStorage(ctrl)
		gomock.InOrder(
			s.EXPECT().FindSlot(uint(1)).Return(&model.Slot{Id: 1, Desc: "Test 1"}, nil),
			s.EXPECT().FindBanner(uint(51)).Return(nil, errors.New("no rows")),
		)

		req, err := http.NewRequest(
			http.MethodPost,
			"/banner/hit",
			strings.NewReader(`{"slot_id":1,"banner_id":51}`),
		)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		HitBannerRequest(s, rr, req)

		expectedResult := "{\"error\":\"no rows\"}\n"
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expectedResult, rr.Body.String())
	})
}
