package service

import (
	"errors"
	"uala-tweet/internal/domain"
	"uala-tweet/internal/dto"
	"uala-tweet/internal/repository/interfaces"
)

var (
	errUserNotFound = errors.New("user not found")
)

type UserService struct {
	repository interfaces.UalaRepository
}

func NewUserService(repository interfaces.UalaRepository) *UserService {
	return &UserService{repository: repository}
}

func (u *UserService) Follow(userID, followID string) error {
	return u.repository.Follow(userID, followID)
}

func (u *UserService) GetTimeline(userID string) ([]dto.TweetDTO, error) {
	user := u.repository.GetUser(userID)
	var tweets []dto.TweetDTO
	if user == nil {
		return nil, errUserNotFound
	}

	tweets = convertTweetsToDTO(u.repository.GetTweets(user))

	return tweets, nil
}

func (u *UserService) GetUser(id string) *dto.UserDTO {
	user := u.convertUserToDTO(id)

	return &user
}

func (u *UserService) GetTweets(userDTO *dto.UserDTO) []dto.TweetDTO {
	user := convertDTOToUser(userDTO)
	tweets := u.repository.GetTweets(&user)

	return convertTweetsToDTO(tweets)
}

func convertDTOToUser(userDTO *dto.UserDTO) domain.User {
	user := domain.User{
		ID:      userDTO.ID,
		Follows: userDTO.Follows,
	}
	return user
}

func (u *UserService) convertUserToDTO(id string) dto.UserDTO {
	user := u.repository.GetUser(id)

	return dto.UserDTO{
		ID:      user.ID,
		Follows: user.Follows,
	}
}

func convertTweetsToDTO(tweets []domain.Tweet) []dto.TweetDTO {
	tweetDTOs := make([]dto.TweetDTO, 0)
	for _, tweet := range tweets {
		tweetDTOs = append(tweetDTOs, convertTweetToDTO(tweet))
	}

	return tweetDTOs
}

func convertTweetToDTO(tweet domain.Tweet) dto.TweetDTO {
	return dto.TweetDTO{
		UserID:  tweet.UserID,
		Content: tweet.Content,
	}
}
