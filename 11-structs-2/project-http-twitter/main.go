package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var id int

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type Id struct {
	Id int `json:"ID"`
}

func main() {
	http.HandleFunc("/tweets", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
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

	id++
	idResp := Id{}
	idResp.Id = id

	resp, err := json.Marshal(idResp)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
	fmt.Printf("Tweet: `%s` from %s", tweet.Message, tweet.Location)
}
