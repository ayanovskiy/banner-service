package handlers

import (
	"banner-service/internal/bandit"
	"banner-service/internal/model"
	"banner-service/internal/server/handlers"
	"banner-service/internal/storage"
	"encoding/json"
	"net/http"
)

type selectBannerRequest struct {
	SlotId  uint `json:"slot_id"`
	GroupId uint `json:"group_id"`
}

type selectBannerResponse struct {
	BannerId uint `json:"banner_id"`
}

// SelectBannerHandler выбирает баннер для показа (многорукий бандит)
func SelectBannerHandler(s storage.IStorage, w http.ResponseWriter, r *http.Request) {
	var req selectBannerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handlers.SendError(w, err)
		return
	}

	group, err := s.FindGroup(req.GroupId)
	if err != nil {
		handlers.SendError(w, err)
		return
	}

	banners, err := s.FindBannersBySlot(req.SlotId)
	if err != nil {
		handlers.SendError(w, err)
		return
	}

	arms := make(bandit.Arms, 0, len(banners))
	stats := make(model.Stats, 0)
	for _, banner := range banners {
		stat, err := s.FindOrCreateStat(req.SlotId, banner.Id, group.Id)
		if err != nil {
			handlers.SendError(w, err)
			return
		}

		stats = append(stats, stat)
		arms = append(arms, bandit.Arm{Count: stat.Shows, Reward: stat.Hits})
	}

	idx := bandit.MultiArmBandit(arms)
	err = s.UpdateStat(stats[idx].Id, stats[idx].Shows+1, stats[idx].Hits)
	if err != nil {
		handlers.SendError(w, err)
		return
	}

	handlers.SendResponse(w, selectBannerResponse{stats[idx].BannerId})
}
