package auth

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

type Config struct {
	SecretKey string `json:"secret_key"`
}

func CreateSession() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config:", err)
		return
	}
	Store = sessions.NewCookieStore([]byte(config.SecretKey))

	// Configure the session store options
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // session expiration time in seconds
		HttpOnly: true,
		Secure:   true, // set to true if using HTTPS
	}
}

func GetSession() *sessions.CookieStore {
	return Store
}
