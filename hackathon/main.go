package main

import (
	"hackathon/controller"
	"hackathon/dao"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch r.URL.Path {
	case "/post":
		switch r.Method {
		case http.MethodPost:
			controller.HandlePost(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	case "/pasttweet":
		switch r.Method {
		case http.MethodPost:
			controller.HandlePastTweet(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	case "/tweetlist":
		switch r.Method {
		case http.MethodPost:
			controller.HandleTweetList(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	case "/replytweet":
		switch r.Method {
		case http.MethodPost:
			controller.HandleReplyTweetlist(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	case "/register":
		switch r.Method {
		case http.MethodPost:
			controller.UserRegister(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	case "/confirmValuationType":
		switch r.Method {
		case http.MethodPost:
			controller.ConfirmValuationType(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	case "/valuation":
		switch r.Method {
		case http.MethodPost:
			controller.HandleValuation(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "無効なエンドポイント", http.StatusNotFound)
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	http.HandleFunc("/", handler)
	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.CloseDBWithSysCall()
	//  コミットする用
	// 8080番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
