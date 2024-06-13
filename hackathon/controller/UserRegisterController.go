package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"net/http"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	var NewUser model.NewUser
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		http.Error(w, "Invalid request body in Register", http.StatusBadRequest)
		return
	}
	if NewUser.UserName == "" {
		http.Error(w, "NoName is not allowed", http.StatusBadRequest)
		return
	}

	err = usecase.UserRegister(NewUser.Email, NewUser.UserName)
	if err != nil {
		http.Error(w, "failed to Register", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User register successfully"))
}
