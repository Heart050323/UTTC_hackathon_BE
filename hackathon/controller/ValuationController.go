package controller

import (
	"encoding/json"
	"fmt"
	"hackathon/model"
	"hackathon/usecase"
	"log"
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

func ConfirmValuationType(w http.ResponseWriter, r *http.Request) {
	var ValuationTypeRequest model.ValuationTypeRequest
	err := json.NewDecoder(r.Body).Decode(&ValuationTypeRequest)
	if err != nil {
		http.Error(w, "Invalid request body in ValuationTypeRequest", http.StatusBadRequest)
		return
	}
	valuationType, err := usecase.GetValuationType(ValuationTypeRequest.TweetID, ValuationTypeRequest.SenderUserID)
	if err != nil {
		http.Error(w, "failed to ValuationTypeRequest", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("ValuationTypeRequest successfully")
	err = json.NewEncoder(w).Encode(valuationType)
	if err != nil {
		http.Error(w, "Failed to encode response valuationType", http.StatusInternalServerError)
		return
	}
	fmt.Println("valuationTypeの呼び出しに成功")
}
