package auth

import "github.com/gorilla/sessions"

var Store *sessions.CookieStore

func CreateSession() {
	Store = sessions.NewCookieStore([]byte("your-secret-key"))

	// Configure the session store options
	Store.Options = &sessions.Options{
		Path:     "/add-user",
		MaxAge:   180, // session expiration time in seconds
		HttpOnly: true,
		Secure:   false, // set to true if using HTTPS
	}
}

func GetSession() *sessions.CookieStore {
	return Store
}
