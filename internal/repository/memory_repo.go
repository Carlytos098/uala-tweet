package repository

import (
	"github.com/google/uuid"
	"sync"

	"uala-tweet/internal/domain"
)

type MemoryRepository struct {
	users  map[string]*domain.User
	tweets map[string]*domain.Tweet
	lock   sync.RWMutex
}

func NewMemoryRepo() *MemoryRepository {
	return &MemoryRepository{
		users:  make(map[string]*domain.User),
		tweets: make(map[string]*domain.Tweet),
	}
}

func (r *MemoryRepository) CreateUser(id string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.users[id] = &domain.User{ID: id}
}

func (r *MemoryRepository) Follow(userID, followID string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if user, exists := r.users[userID]; exists {
		user.Follows = append(user.Follows, followID)
	}

	return nil
}

func (r *MemoryRepository) CreateTweet(tweet *domain.Tweet) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	tweet.ID = uuid.New().String()
	r.tweets[tweet.ID] = tweet

	return nil
}

func (r *MemoryRepository) GetTweets(user *domain.User) []domain.Tweet {
	r.lock.RLock()
	defer r.lock.RUnlock()
	var result []domain.Tweet
	for _, followID := range user.Follows {
		for _, tweet := range r.tweets {
			if tweet.UserID == followID {
				result = append(result, *tweet)
			}
		}
	}
	return result
}

func (r *MemoryRepository) GetAllTweets() map[string]*domain.Tweet {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.tweets
}

func (r *MemoryRepository) GetUser(id string) *domain.User {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.users[id]
}
