package models

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type ID struct {
	Id int `json:"ID"`
}

type TweetsList struct {
	Tweets []Tweet `json:"tweets"`
}
