package controller

import (
	"encoding/json"
	"net/http"
	"uala-tweet/internal/dto"
)

type UserUseCase interface {
	Follow(userID string, followID string)
	GetTimeline(userID string) []dto.TweetDTO
}

type UserController struct {
	useCase UserUseCase
}

func NewUserController(useCase UserUseCase) *UserController {
	return &UserController{useCase: useCase}
}

func (c *UserController) Follow(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")
	followID := r.URL.Query().Get("followID")
	c.useCase.Follow(userID, followID)
	w.WriteHeader(http.StatusNoContent)
}

func (c *UserController) GetTimeline(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")
	tweets := c.useCase.GetTimeline(userID)
	err := json.NewEncoder(w).Encode(tweets)
	if err != nil {
		return
	}
}
