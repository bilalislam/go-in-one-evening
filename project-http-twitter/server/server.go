package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"twitter/models"
	"twitter/repository"
)

type Server struct {
	Repository repository.TweetRepository
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tweet := models.Tweet{}
	if err := json.Unmarshal(body, &tweet); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := s.Repository.AddTweet(tweet)
	if err != nil {
		log.Println("Failed to add tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	idResp := models.ID{
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

func (s Server) ListTweets(w http.ResponseWriter, r *http.Request) {
	getTweetList, _ := s.Repository.Tweets()
	tweetList := models.TweetsList{
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
