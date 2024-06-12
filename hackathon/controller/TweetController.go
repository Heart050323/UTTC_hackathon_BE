package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"net/http"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	var PostTweet model.PostTweet
	err := json.NewDecoder(r.Body).Decode(&PostTweet)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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
