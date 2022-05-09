package controller

import (
	"encoding/json"
	"fmt"
	"github.com/CodaGott/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

type UserController struct {
	session *mgo.Session
}

func (userController UserController)GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	objectId := bson.ObjectIdHex(id)

	user := models.User{}

	if err := userController.session.DB("mongo-golang").C("users").FindId(objectId).One(&user); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", uj)
}

func (userController UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}

	json.NewDecoder(r.Body).Decode(&user)

	user.Id = bson.NewObjectId()

	userController.session.DB("mongo-golang").C("users").Insert(user)

	uj, err := json.Marshal(user)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (userController UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	objectId := bson.ObjectIdHex(id)

	if err := userController.session.DB("mongo-golang").C("users").RemoveId(objectId); err != nil{
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "Deleted user", objectId, "\n")
}