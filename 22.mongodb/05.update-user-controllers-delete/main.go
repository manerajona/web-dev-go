package main

import (
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manerajona/web-dev-go/22.mongodb/05.update-user-controllers-delete/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

// docker container run -p 27017:27017 --name mongo -d mongo
