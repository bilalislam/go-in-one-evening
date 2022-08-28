package repository

import (
	"errors"
	"twitter/models"
)

type TweetRepository interface {
	AddTweet(t models.Tweet) (int, error)
	Tweets() ([]models.Tweet, error)
}

type TweetMemoryRepository struct {
	tweets []models.Tweet
}

func (t *TweetMemoryRepository) AddTweet(tw models.Tweet) (int, error) {
	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

func (t TweetMemoryRepository) Tweets() ([]models.Tweet, error) {
	if len(t.tweets) > 0 {
		return t.tweets, nil
	}

	return nil, errors.New("tweets not found !")
}
