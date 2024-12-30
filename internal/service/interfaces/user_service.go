package interfaces

import (
	"uala-tweet/internal/dto"
)

type UserService interface {
	Follow(userID, followID string) error
	GetTimeline(userID string) ([]dto.TweetDTO, error)
	GetUser(id string) *dto.UserDTO
	GetTweets(user *dto.UserDTO) []dto.TweetDTO
}
