package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manerajona/web-dev-go/22.mongodb/01.update-user-controller/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession()) // Get a UserController instance

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

// docker container run -p 27017:27017 --name mongo -d mongo
