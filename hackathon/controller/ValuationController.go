package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"net/http"
)

func HandleValuation(w http.ResponseWriter, r *http.Request) {
	var Valuation model.Valuation
	err := json.NewDecoder(r.Body).Decode(&Valuation)
	if err != nil {
		http.Error(w, "Invalid request body in HandleValuation", http.StatusBadRequest)
		return
	}
	err = usecase.Valuation(Valuation.TweetID, Valuation.SenderUserID, Valuation.ValuationType)
	if err != nil {
		http.Error(w, "failed to Valuation", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("VAluation successfully"))
}
