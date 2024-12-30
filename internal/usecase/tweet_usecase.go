package usecase

import (
	"uala-tweet/internal/dto"
	"uala-tweet/internal/service/interfaces"
)

type TweetUseCase struct {
	service interfaces.TweetService
}

func NewTweetUseCase(tweetService interfaces.TweetService) *TweetUseCase {
	return &TweetUseCase{service: tweetService}
}

func (uc *TweetUseCase) PublishTweet(dto dto.TweetDTO) error {
	err := uc.service.PublishTweet(dto)
	if err != nil {
		return err
	}

	return nil
}
