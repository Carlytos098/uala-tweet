package config

import (
	"net/http"
	"uala-tweet/internal/controller"
	"uala-tweet/internal/middleware"
)

func SetupRoutes(userController *controller.UserController, tweetController *controller.TweetController) {
	http.HandleFunc("/tweets", middleware.AuthMiddleware(tweetController.PublishTweet))
	http.HandleFunc("/follow", middleware.AuthMiddleware(userController.Follow))
	http.HandleFunc("/timeline", middleware.AuthMiddleware(userController.GetTimeline))
}
