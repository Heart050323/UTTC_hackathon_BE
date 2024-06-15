package controller

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

func HandleTweetList(w http.ResponseWriter, r *http.Request) {
	tweetList, err := usecase.TweetList()
	if err != nil {
		http.Error(w, "failed to TweetList", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("tweetListCall successfully")
	err = json.NewEncoder(w).Encode(tweetList)
	if err != nil {
		http.Error(w, "Failed to encode response tweetlist", http.StatusInternalServerError)
		return
	}
	fmt.Println(tweetList)

}

func HandleReplyTweetlist(w http.ResponseWriter, r *http.Request) {
	var ReplyRequest model.ReplyRequest
	err := json.NewDecoder(r.Body).Decode(&ReplyRequest)
	if err != nil {
		http.Error(w, "Invalid request body in ReplyRequest", http.StatusBadRequest)
		return
	}
	replyTweetList, err := usecase.ReplyTweetList(ReplyRequest.RepliedTweetID)
	if err != nil {
		http.Error(w, "failed to ReplyRequest", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("ReplyRequest successfully")
	err = json.NewEncoder(w).Encode(replyTweetList)
	if err != nil {
		http.Error(w, "Failed to encode response replytweet", http.StatusInternalServerError)
		return
	}
	fmt.Println(replyTweetList)
}

func HandlePastTweet(w http.ResponseWriter, r *http.Request) {
	var EmailRequest model.EmailRequest
	err := json.NewDecoder(r.Body).Decode(&EmailRequest)
	if err != nil {
		http.Error(w, "Invalid request body in EmailRequest", http.StatusBadRequest)
		return
	}
	wholeData, err := usecase.PastTweet(EmailRequest.Email)
	if err != nil {
		http.Error(w, "failed to EmailRequest", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("EmailRequest successfully")
	err = json.NewEncoder(w).Encode(wholeData)
	if err != nil {
		http.Error(w, "Failed to encode response Pasttweet", http.StatusInternalServerError)
		return
	}
	fmt.Println(wholeData)
}
