package main

import (
	"github.com/CodaGott/mongo-golang/controller"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	route := httprouter.New()

	userController := controller.NewUserController(getSession())

	route.GET("/user/:id", userController.GetUser)
	route.POST("/user", userController.CreateUser)
	route.DELETE("/user/:id", userController.DeleteUser)
	http.ListenAndServe("localhost:9000", route)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		log.Fatalln("Error occurred connecting to db")
		panic(err)
	}

	return s
}