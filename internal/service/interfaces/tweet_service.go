package interfaces

import (
	"uala-tweet/internal/dto"
)

type TweetService interface {
	PublishTweet(dto dto.TweetDTO) error
}
