package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	mongoSession := getSession()
	apiCollection := mongoSession.DB("cache").C("api")
	cacheCollection := mongoSession.DB("cache").C("cache")

	apiController := controllers.NewApiController(apiCollection)
	cacheController := controllers.NewCacheController(apiCollection, cacheCollection)
	router := httprouter.New()
	router.GET("/api", apiController.GetApi)
	router.GET("/api/:apikey", cacheController.GetAll)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
