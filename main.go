package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"./controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	mongoSession := getSession()
	apiCollection := mongoSession.DB("cache").C("api")
	apiController := controllers.NewApiController(apiCollection)
	router := httprouter.New()
	router.GET("/api", apiController.GetApi)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
