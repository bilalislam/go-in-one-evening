package repository

import (
	"errors"
	"sync"
	"twitter/models"
)

type TweetRepository interface {
	AddTweet(t models.Tweet) (int, error)
	Tweets() ([]models.Tweet, error)
}

type TweetMemoryRepository struct {
	lock   sync.RWMutex
	tweets []models.Tweet
}

func (t *TweetMemoryRepository) AddTweet(tw models.Tweet) (int, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.tweets = append(t.tweets, tw)
	return len(t.tweets), nil
}

func (t *TweetMemoryRepository) Tweets() ([]models.Tweet, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if len(t.tweets) > 0 {
		return t.tweets, nil
	}

	return nil, errors.New("tweets not found")
}
