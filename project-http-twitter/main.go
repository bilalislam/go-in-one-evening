package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"twitter/models"
	"twitter/repository"
	"twitter/server"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

func main() {
	s := server.Server{
		Repository: &repository.TweetMemoryRepository{},
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.With(httprate.LimitByIP(10, 1*time.Minute)).Post("/tweets", s.AddTweet)
	r.Get("/tweets", s.ListTweets)

	go spamTweets()

	log.Fatal(http.ListenAndServe(":8080", r))
}

func spamTweets() error {
	for {
		addTweetPayload := models.Tweet{
			Message:  "ass",
			Location: "ass",
		}

		marshaledPayload, err := json.Marshal(addTweetPayload)
		if err != nil {
			fmt.Print(err)
			return err
		}
		url := "http://localhost:8080/tweets"

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if err != nil {
			fmt.Print(err)
			return err
		}
		defer resp.Body.Close()

		if err != nil {
			return err
		}

		id := models.ID{}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
			return err
		}

		json.Unmarshal(body, &id)
		_, err = fmt.Printf("added tweet with : %d ", id)
		if err != nil {
			fmt.Print(err)
			return err
		}
	}
}
