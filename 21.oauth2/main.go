package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
	<body>
		<form action="/oauth2/github" method="post">
			<input type="submit" value="Login with GitHub">
		</form>
	</body>
	</html>`)
}

func auth(w http.ResponseWriter, r *http.Request) {
	// state is a token to protect the user from CSRF attacks
	state := uuid.New().String()

	redirectURL := GitHubOauthConfig.AuthCodeURL(state)

	http.SetCookie(w, &http.Cookie{
		Name:  CookieName,
		Value: state,
	})
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func authCallback(w http.ResponseWriter, r *http.Request) {

	if cookie, _ := r.Cookie(CookieName); cookie != nil && cookie.Value != "" {

		state := r.FormValue("state")
		if state != cookie.Value {
			http.Error(w, "invalid state", http.StatusBadRequest)
			return
		}

		code := r.FormValue("code")
		token, err := GitHubOauthConfig.Exchange(r.Context(), code)
		if err != nil {
			http.Error(w, "Couldn't login", http.StatusInternalServerError)
			return
		}

		gitHubUsr, err := fetchUserData(r.Context(), token)
		if err != nil {
			http.Error(w, "Couldn't obtain user data", http.StatusInternalServerError)
			return
		}

		io.WriteString(w, `<!DOCTYPE html><html><body><h1>Welcome `+gitHubUsr.Data.Viewer.Name+`!</h1></body></html>`)
		return
	}
	io.WriteString(w, `<!DOCTYPE html><html><body><h1>Access denied</h1><a href="/login">sign in</a></body></html>`)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/oauth2/github", auth)
	http.HandleFunc("/oauth2/receive", authCallback)
	http.ListenAndServe(":8080", nil)
}
