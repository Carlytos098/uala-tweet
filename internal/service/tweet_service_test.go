package service_test

import (
	"errors"
	"testing"
	"uala-tweet/internal/domain"
	"uala-tweet/internal/repository/interfaces"
	"uala-tweet/internal/service"

	"uala-tweet/internal/dto"
)

type mockRepository struct {
	mockCreateUser   func(id string)
	mockFollow       func(userID, followID string) error
	mockCreateTweet  func(tweet *domain.Tweet) error
	mockGetTweets    func(user *domain.User) []domain.Tweet
	mockGetAllTweets func() map[string]*domain.Tweet
	mockGetUser      func(id string) *domain.User
}

func (m mockRepository) CreateUser(id string) {
	m.mockCreateUser(id)
}

func (m mockRepository) Follow(userID, followID string) error {
	return m.mockFollow(userID, followID)
}

func (m mockRepository) CreateTweet(tweet *domain.Tweet) error {
	return m.mockCreateTweet(tweet)
}

func (m mockRepository) GetTweets(user *domain.User) []domain.Tweet {
	return m.mockGetTweets(user)
}

func (m mockRepository) GetAllTweets() map[string]*domain.Tweet {
	return m.mockGetAllTweets()
}

func (m mockRepository) GetUser(id string) *domain.User {
	return m.mockGetUser(id)
}

func TestTweetService_PublishTweet(t1 *testing.T) {
	type fields struct {
		repository interfaces.UalaRepository
	}
	type args struct {
		dto dto.TweetDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				repository: mockRepository{
					mockCreateTweet: func(tweet *domain.Tweet) error {
						return nil
					},
				},
			},
			args: args{dto: dto.TweetDTO{
				UserID:  "User1",
				Content: "Content...",
			}},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				repository: mockRepository{
					mockCreateTweet: func(tweet *domain.Tweet) error {
						return errors.New("some error")
					},
				},
			},
			args: args{dto: dto.TweetDTO{
				UserID:  "User1",
				Content: "Content...",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			s := service.NewTweetService(tt.fields.repository)
			if err := s.PublishTweet(tt.args.dto); (err != nil) != tt.wantErr {
				t1.Errorf("PublishTweet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

	}
}
