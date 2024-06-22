package controller

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

func HandleUserInfo(w http.ResponseWriter, r *http.Request) {
	var EmailRequest model.EmailRequest
	err := json.NewDecoder(r.Body).Decode(&EmailRequest)
	if err != nil {
		http.Error(w, "Invalid request body in EmailRequest", http.StatusBadRequest)
		return
	}
	userData, err := usecase.GetUserInfo(EmailRequest.Email)
	if err != nil {
		http.Error(w, "failed to EmailRequest", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("EmailRequest successfully")
	err = json.NewEncoder(w).Encode(userData)
	if err != nil {
		http.Error(w, "Failed to encode response userData", http.StatusInternalServerError)
		return
	}
	fmt.Println(userData)
}

func HandleUserProfile(w http.ResponseWriter, r *http.Request) {
	var UserIDRequest model.UserIDRequest
	err := json.NewDecoder(r.Body).Decode(&UserIDRequest)
	if err != nil {
		http.Error(w, "Invalid request body in UserIDRequest", http.StatusBadRequest)
		return
	}
	userProfile, err := usecase.GetUserProfile(UserIDRequest.User_id)
	if err != nil {
		http.Error(w, "failed to UserIDRequest", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("UserIDRequest successfully")
	err = json.NewEncoder(w).Encode(userProfile)
	if err != nil {
		http.Error(w, "Failed to encode response userProfile", http.StatusInternalServerError)
		return
	}
	fmt.Println(userProfile)
}

func HandleUserProfileModify(w http.ResponseWriter, r *http.Request) {
	var newUserProfile model.UserProfile
	err := json.NewDecoder(r.Body).Decode(&newUserProfile)
	if err != nil {
		http.Error(w, "Invalid request body in newUserProfileRequest", http.StatusBadRequest)
		return
	}
	err = usecase.UserProfileModify(newUserProfile.User_id, newUserProfile.User_Name, newUserProfile.StatusMessage)
	if err != nil {
		http.Error(w, "failed to newUserProfileRequest", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("UserProfileModify successfully")
}
