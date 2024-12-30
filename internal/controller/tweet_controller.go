package controller

import (
	"encoding/json"
	"net/http"

	"uala-tweet/internal/dto"
)

type TweetUseCase interface {
	PublishTweet(dto dto.TweetDTO) error
}

type TweetController struct {
	useCase TweetUseCase
}

func NewTweetController(useCase TweetUseCase) *TweetController {
	return &TweetController{useCase: useCase}
}

func (c *TweetController) PublishTweet(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("User-ID")
	if userID == "" {
		http.Error(w, "User-ID header is required", http.StatusUnauthorized)
		return
	}

	var tweetDTO dto.TweetDTO
	if err := json.NewDecoder(r.Body).Decode(&tweetDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tweetDTO.UserID = userID
	if err := c.useCase.PublishTweet(tweetDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
