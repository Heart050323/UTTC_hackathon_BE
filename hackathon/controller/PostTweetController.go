package controller

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	var PostTweet model.PostTweet
	err := json.NewDecoder(r.Body).Decode(&PostTweet)
	if err != nil {
		http.Error(w, "Invalid request body in Post", http.StatusBadRequest)
		return
	}
	if len(PostTweet.Content) > 140 {
		http.Error(w, "Invalid Tweet", http.StatusBadRequest)
		return
	}
	err = usecase.PostTweet(PostTweet.SenderUserID, PostTweet.Content, PostTweet.RepliedTweetID, PostTweet.ReTweetID)
	if err != nil {
		http.Error(w, "failed to PostTweet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tweet posted successfully"))

}

func HandleRetweetOn(w http.ResponseWriter, r *http.Request) {
	var RetweetOnRequest model.RetweetOnRequest
	err := json.NewDecoder(r.Body).Decode(&RetweetOnRequest)
	if err != nil {
		http.Error(w, "Invalid request body in REtweetOnRequest", http.StatusBadRequest)
		return
	}
	RetweetOn, err := usecase.GetRetweetOn(RetweetOnRequest.TweetID, RetweetOnRequest.SenderUserID)
	if err != nil {
		http.Error(w, "failed to RetweetOnRequest", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("RetweetOnRequest successfully")
	err = json.NewEncoder(w).Encode(RetweetOn)
	if err != nil {
		http.Error(w, "Failed to encode response RetweetOn", http.StatusInternalServerError)
		return
	}
	fmt.Println(RetweetOn)
}
