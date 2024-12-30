package service

import (
	"uala-tweet/internal/domain"
	"uala-tweet/internal/dto"
	"uala-tweet/internal/repository/interfaces"
)

type TweetService struct {
	repo interfaces.UalaRepository
}

func NewTweetService(repo interfaces.UalaRepository) *TweetService {
	return &TweetService{repo}
}

func (t *TweetService) PublishTweet(dto dto.TweetDTO) error {
	tweet := domain.Tweet{
		UserID:  dto.UserID,
		Content: dto.Content,
	}

	return t.repo.CreateTweet(&tweet)
}
