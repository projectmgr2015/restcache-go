package controllers

import (
	"../model"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type CacheController struct {
	apiCollection   *mgo.Collection
	cacheCollection *mgo.Collection
}

func NewCacheController(apiCollection *mgo.Collection, cacheCollection *mgo.Collection) *CacheController {
	return &CacheController{apiCollection, cacheCollection}
}

func (ac CacheController) GetAll(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	apiKey := params.ByName("apikey")
	api := model.Api{}

	err := ac.apiCollection.FindId(apiKey).One(&api)
	if err != nil {
		response.WriteHeader(404)
		return
	}

	caches := []model.Cache{}
	ac.cacheCollection.Find(bson.M{"api": apiKey}).All(&caches)

	json, _ := json.Marshal(caches)
	response.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(response, "%s", json)
}
