package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"uala-tweet/internal/controller"
	"uala-tweet/internal/dto"
)

type mockTweetUseCase struct {
	mockPublishTweet func(dto dto.TweetDTO) error
}

func (m mockTweetUseCase) PublishTweet(dto dto.TweetDTO) error {
	return m.mockPublishTweet(dto)
}

func TestTweetController_PublishTweet(t *testing.T) {
	type fields struct {
		useCase controller.TweetUseCase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	dataTest := args{
		w: httptest.NewRecorder(),
		r: &http.Request{
			Header: http.Header{"User-ID": []string{"user1"}},
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
				useCase: &mockTweetUseCase{
					mockPublishTweet: func(dto dto.TweetDTO) error {
						return nil
					},
				},
			},
			args: dataTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := controller.NewTweetController(tt.fields.useCase)
			c.PublishTweet(tt.args.w, tt.args.r)
		})
	}
}
