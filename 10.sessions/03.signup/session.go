package main

import (
	"net/http"
)

func getUser(req *http.Request) (usr user) {
	// get cookie
	cookie, err := req.Cookie("session")
	if err == nil {

		// if the user exists already, get user
		if un, ok := dbSessions[cookie.Value]; ok {
			usr = dbUsers[un]
		}
	}
	return
}

func alreadyLoggedIn(req *http.Request) (ok bool) {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok = dbUsers[un]
	return
}
