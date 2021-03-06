package main

import (
	"./controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	apiCollection := getSession().DB("cache").C("api")
	cacheCollection := getSession().DB("cache").C("cache")

	apiController := controllers.NewApiController(apiCollection)
	cacheController := controllers.NewCacheController(apiCollection, cacheCollection)

	router := httprouter.New()
	router.GET("/api", apiController.GetApi)
	router.GET("/api/:apikey", cacheController.GetAll)
	router.GET("/api/:apikey/:key", cacheController.GetOne)
	router.POST("/api/:apikey/:key", cacheController.Create)
	router.PUT("/api/:apikey/:key", cacheController.Update)
	router.DELETE("/api/:apikey/:key", cacheController.Delete)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://mongo.projectmgr2015.tk")
	if err != nil {
		panic(err)
	}
	return s
}
