package usecase

import (
	"uala-tweet/internal/dto"
	"uala-tweet/internal/service/interfaces"
)

type UserUseCase struct {
	service interfaces.UserService
}

func NewUserUseCase(service interfaces.UserService) *UserUseCase {
	return &UserUseCase{service: service}
}

func (uc *UserUseCase) Follow(userID string, followID string) {
	err := uc.service.Follow(userID, followID)
	if err != nil {
		return
	}
}

func (uc *UserUseCase) GetTimeline(userID string) []dto.TweetDTO {
	user := uc.service.GetUser(userID)
	var tweets []dto.TweetDTO
	if user != nil {
		tweets = uc.service.GetTweets(user)
	}

	return tweets
}
