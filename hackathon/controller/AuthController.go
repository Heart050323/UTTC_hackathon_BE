package controller

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
	"hackathon/usecase"
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Auth successfully"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
	fmt.Println(userInfo)
}
