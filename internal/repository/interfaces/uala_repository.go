package interfaces

import "uala-tweet/internal/domain"

type UalaRepository interface {
	CreateUser(id string)
	Follow(userID, followID string) error
	CreateTweet(tweet *domain.Tweet) error
	GetTweets(user *domain.User) []domain.Tweet
	GetAllTweets() map[string]*domain.Tweet
	GetUser(id string) *domain.User
}
