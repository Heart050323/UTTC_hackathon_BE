package controller

import (
	"encoding/json"
	"kaizen/dao"
	"kaizen/model"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid"
)

func handlePost(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if user.Name == "" || len(user.Name) > 50 || user.Age < 20 || user.Age > 80 {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	id := generateULID()

	err = dao.RegisterUser(id, user.Name, user.Age)
	if err != nil {
		http.Error(w, "failed to RegisterUser", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{ID: id})
}
func generateULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
