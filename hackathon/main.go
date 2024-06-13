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

	switch r.URL.Path {
	case "/post":
		switch r.Method {
		case http.MethodPost:
			controller.HandlePost(w, r)
		default:
			http.Error(w, "許可されていないメソッド", http.StatusMethodNotAllowed)
		}
	case "/auth":
		switch r.Method {
		case http.MethodPost:
			controller.HandleAuth(w, r)
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
