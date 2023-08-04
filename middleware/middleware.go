package middleware

import (
	"fmt"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle, store *sessions.CookieStore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session, err := store.Get(r, "login_session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			fmt.Println("Session error")
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusFound)
			fmt.Println(auth)
			fmt.Println(ok)
			return
		}

		// if session.IsNew {
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
		// 	fmt.Println("Session habis 2")
		// } else {
		// 	fmt.Println("session valid")
		// 	fmt.Println(session)
		// 	fmt.Println(session.IsNew)
		// 	tmpl, err := template.ParseFiles("templates/user-form.html")

		// 	users := userservice.GetUser()
		// 	fmt.Println(users)

		// 	err = tmpl.Execute(w, users)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// }
		next(w, r, ps)
	}
}
