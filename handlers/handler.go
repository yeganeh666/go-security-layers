package handlers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"security-example/entity"
	"security-example/pkg/auth"
)

var oauthConfig = oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{"profile", "email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://provider.com/auth",
		TokenURL: "https://provider.com/token",
	},
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	storedPassword, ok := entity.Users[user.Username]
	if !ok || storedPassword != user.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateToken(user.Username)
	if err != nil {
		http.Error(w, fmt.Sprint("Internal Server Error ", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}

func Data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"data": "%s"}`, "Sensitive Data")
}

func OAuth(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, oauthConfig.AuthCodeURL("state"), http.StatusTemporaryRedirect)
}

func OAuthCallback(w http.ResponseWriter, r *http.Request) {
	//code := r.URL.Query().Get("code")
	//token, err := oauthConfig.Exchange(r.Context(), code)
	//if err != nil {
	//	http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
	//	return
	//}

	// Get user information using token.AccessToken
}
