package main

import (
	"log"
	"net/http"
	"security-example/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/data", handlers.AuthMiddleware(handlers.Data))
	http.HandleFunc("/oauth", handlers.OAuth)
	http.HandleFunc("/callback", handlers.OAuthCallback)

	log.Println("app is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
