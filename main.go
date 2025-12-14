package main

import (
	"chat-room-app/handlers"
	"chat-room-app/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	services.InitPusher()

	router := mux.NewRouter()

	router.Use(corsMiddleware)

	router.HandleFunc("/create-room", handlers.CreateRoom).Methods("POST", "OPTIONS")
	router.HandleFunc("/send-message", handlers.SendMessage).Methods("POST", "OPTIONS")
	router.HandleFunc("/create-room", handlers.JoinRoom).Methods("POST", "OPTIONS")

	log.Println("Server Starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
