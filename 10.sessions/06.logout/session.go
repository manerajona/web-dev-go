package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) (usr user) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}

	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	if un, ok := dbSessions[c.Value]; ok {
		usr = dbUsers[un]
	}
	return
}

func alreadyLoggedIn(req *http.Request) (ok bool) {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok = dbUsers[un]
	return
}
