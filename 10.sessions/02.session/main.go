package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
		http.SetCookie(w, cookie)
	}

	// if the user exists already, get user
	var usr user
	if userName, ok := dbSessions[cookie.Value]; ok {
		usr = dbUsers[userName]
	}

	// process form submission
	if req.Method == http.MethodPost {
		userName := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		usr = user{userName, fn, ln}
		dbSessions[cookie.Value] = userName
		dbUsers[userName] = usr
	}

	tpl.ExecuteTemplate(w, "index.gohtml", usr)
}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	userName, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	usr := dbUsers[userName]
	tpl.ExecuteTemplate(w, "bar.gohtml", usr)
}
