package service_test

import (
	"errors"
	"reflect"
	"testing"
	"uala-tweet/internal/domain"
	"uala-tweet/internal/dto"
	"uala-tweet/internal/repository/interfaces"
	"uala-tweet/internal/service"
)

func TestUserService_Follow(t *testing.T) {
	type fields struct {
		repository interfaces.UalaRepository
	}
	type args struct {
		userID   string
		followID string
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
					mockFollow: func(userID string, followID string) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				repository: mockRepository{
					mockFollow: func(userID string, followID string) error {
						return errors.New("some error")
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.NewUserService(tt.fields.repository)
			if err := s.Follow(tt.args.userID, tt.args.followID); (err != nil) != tt.wantErr {
				t.Errorf("Follow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_GetTimeline(t *testing.T) {
	type fields struct {
		repository interfaces.UalaRepository
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.TweetDTO
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				repository: mockRepository{
					mockGetUser: func(id string) *domain.User {
						return &domain.User{}
					},
					mockGetTweets: func(user *domain.User) []domain.Tweet {
						return []domain.Tweet{}
					},
				},
			},
			want:    make([]dto.TweetDTO, 0),
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				repository: mockRepository{
					mockGetUser: func(id string) *domain.User {
						return nil
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.NewUserService(tt.fields.repository)
			if got, err := s.GetTimeline(tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTimeline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
