package controller

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
	"hackathon/usecase"
	"log"
	"net/http"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var Auth model.UserAuth
	err := json.NewDecoder(r.Body).Decode(&Auth)
	if err != nil {
		http.Error(w, "Invalid request body in Auth", http.StatusBadRequest)
		return
	}
	userInfo, err := usecase.Auth(Auth.Email)
	if err != nil {
		http.Error(w, "failed to Auth", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("Auth successfully")
	err = json.NewEncoder(w).Encode(userInfo)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	fmt.Println(userInfo)
}
