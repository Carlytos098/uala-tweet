package main

import (
	"log"
	"net/http"
	"uala-tweet/internal/config"
	"uala-tweet/internal/controller"
	"uala-tweet/internal/domain"
	"uala-tweet/internal/factory"
	"uala-tweet/internal/service"
	"uala-tweet/internal/usecase"
)

func main() {
	repoFactory := factory.NewFactory()
	repo := repoFactory.NewMemoryRepository()

	repo.CreateUser("user1")
	repo.CreateUser("user2")
	repo.CreateUser("user3")
	_ = repo.Follow("user1", "user2")
	err := repo.CreateTweet(&domain.Tweet{
		ID:      "1",
		UserID:  "user2",
		Content: "Este es un tweet del user2.",
	})
	if err != nil {
		return
	}

	tweetService := service.NewTweetService(repo)
	userService := service.NewUserService(repo)

	tweetUseCase := usecase.NewTweetUseCase(tweetService)
	userUseCase := usecase.NewUserUseCase(userService)

	tweetController := controller.NewTweetController(tweetUseCase)
	userController := controller.NewUserController(userUseCase)

	config.SetupRoutes(userController, tweetController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
