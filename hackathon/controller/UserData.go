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
