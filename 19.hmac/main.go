package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, req *http.Request) {

	cookie, _ := req.Cookie(CookieName)
	if cookie == nil || cookie.Value == "" {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	// decode and extract username and signature
	decodedSessionId, _ := base64.StdEncoding.DecodeString(cookie.Value)

	if pair := strings.Split(string(decodedSessionId), ":"); len(pair) == 2 {
		username := pair[0]
		signature := pair[1]

		// validate signature
		if checkSignature(signature, FixedPassword) {
			io.WriteString(w, `<!DOCTYPE html><html><body><h1>Welcome `+username+`!</h1><a href="/logout">sign off</a></body></html>`)
			return
		}
	}

	cookie.Value = ""
	http.SetCookie(w, cookie)
	io.WriteString(w, `<!DOCTYPE html><html><body><h1>Access denied</h1><a href="/login">sign in</a></body></html>`)
}

func login(w http.ResponseWriter, req *http.Request) {

	cookie, _ := req.Cookie(CookieName)
	if cookie != nil && cookie.Value != "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		// hmac signature
		signature := signWithSha512(password)
		fmt.Printf("signature=%x\n", signature)

		// encode pair username:signature in base64 as sessionId
		sessionId := base64.StdEncoding.EncodeToString([]byte(username + ":" + signature))
		fmt.Printf("%v=%x\n", CookieName, sessionId)

		http.SetCookie(w, &http.Cookie{
			Name:  CookieName,
			Value: sessionId,
		})
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	<body>
		<form method="POST">
			<label>username : </label>
	      	<input type="email" placeholder="Enter email" name="username" required>
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
