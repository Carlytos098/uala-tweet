package controller_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"uala-tweet/internal/controller"
	"uala-tweet/internal/dto"
)

type mockUserUseCase struct {
	mockFollow      func(userID string, followID string)
	mockGetTimeline func(userID string) []dto.TweetDTO
}

func (m mockUserUseCase) Follow(userID, followID string) {
	m.mockFollow(userID, followID)
}

func (m mockUserUseCase) GetTimeline(userID string) []dto.TweetDTO {
	return m.mockGetTimeline(userID)
}

func TestUserController_Follow(t *testing.T) {
	type fields struct {
		useCase controller.UserUseCase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	dataTest := args{
		w: httptest.NewRecorder(),
		r: &http.Request{
			Header: http.Header{"User-ID": []string{"user1"}},
			URL:    &url.URL{RawQuery: "followID=user2"},
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			fields: fields{
				useCase: &mockUserUseCase{
					mockFollow: func(userID string, followID string) {},
				},
			},
			args: dataTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := controller.NewUserController(tt.fields.useCase)
			c.Follow(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_GetTimeline(t *testing.T) {
	type fields struct {
		useCase controller.UserUseCase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	dataTest := args{
		w: httptest.NewRecorder(),
		r: &http.Request{
			Header: http.Header{"User-ID": []string{"user1"}},
			URL:    &url.URL{RawQuery: "followID=user2"},
		},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			fields: fields{
				useCase: mockUserUseCase{
					mockGetTimeline: func(userID string) []dto.TweetDTO {
						return []dto.TweetDTO{}
					},
				},
			},
			args: dataTest,
		},
		{
			name: "fail/no-tweets",
			fields: fields{
				useCase: mockUserUseCase{
					mockGetTimeline: func(userID string) []dto.TweetDTO {
						return nil
					},
				},
			},
			args: dataTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := controller.NewUserController(tt.fields.useCase)
			c.GetTimeline(tt.args.w, tt.args.r)
			println(tt.args.w)
		})
	}
}
