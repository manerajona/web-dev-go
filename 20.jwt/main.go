package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(CookieName)
	if cookie != nil && cookie.Value != "" {
		claims, err := parseToken(cookie.Value)
		if err == nil {
			io.WriteString(w, `<!DOCTYPE html><html><body><h1>Welcome `+claims.Username+`!</h1><a href="/logout">sign off</a></body></html>`)
			return
		}
		log.Println(err)
	}
	io.WriteString(w, `<!DOCTYPE html><html><body><h1>Access denied</h1><a href="/login">sign in</a></body></html>`)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("Username")
		//password := r.FormValue("password") // not used

		token, err := createToken(email)
		if err != nil {
			http.Error(w, "couldn't create access token", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		fmt.Printf("%v=%x\n", CookieName, token)

		http.SetCookie(w, &http.Cookie{
			Name:  CookieName,
			Value: token,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	<body>
		<form method="POST">
			<label>Username : </label>
	      	<input type="email" placeholder="Enter email" name="Username" required>
            <label>Password : </label>   
            <input type="password" placeholder="Enter Password" name="password" required>
	      	<input type="submit">
	    </form>
	</body>
	</html>`)
}

func logout(w http.ResponseWriter, req *http.Request) {
	cookie := &http.Cookie{
		Name:  CookieName,
		Value: "",
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}
