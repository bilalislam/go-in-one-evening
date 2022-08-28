package main

import (
	"encoding/json"
	errors "errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TweetRepository interface {
	AddTweet(t Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetMemoryRepository struct {
	tweets []Tweet
}

func (t *TweetMemoryRepository) AddTweet(tw Tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

func (t TweetMemoryRepository) Tweets() ([]Tweet, error) {
	if len(t.tweets) > 0 {
		return t.tweets, nil
	}

	return nil, errors.New("tweets not found !")
}

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type ID struct {
	Id int `json:"ID"`
}

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}

type server struct {
	repository TweetRepository
}

func (s server) tweets(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer time.Since(start)
	duration := start

	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listTweets(w, r)
	}
	fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
}

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tweet := Tweet{}
	if err := json.Unmarshal(body, &tweet); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.repository.AddTweet(tweet)
	if err != nil {
		log.Println("Failed to add tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	idResp := ID{
		Id: id,
	}

	resp, err := json.Marshal(idResp)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {
	getTweetList, _ := s.repository.Tweets()
	tweetList := tweetsList{
		Tweets: getTweetList,
	}

	resp, err := json.Marshal(tweetList)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func main() {
	s := server{
		repository: &TweetMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
