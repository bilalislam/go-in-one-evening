package main

import (
	"log"
	"net/http"
	"twitter/repository"
	"twitter/server"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	s := server.Server{
		Repository: &repository.TweetMemoryRepository{},
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Post("/tweets", s.AddTweet)
	r.Get("/tweets", s.ListTweets)

	log.Fatal(http.ListenAndServe(":8080", r))
}
