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

func (ac CacheController) GetOne(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	apiKey := params.ByName("apikey")
	key := params.ByName("key")
	api := model.Api{}

	apiError := ac.apiCollection.FindId(apiKey).One(&api)
	if apiError != nil {
		response.WriteHeader(404)
		return
	}

	cache := model.Cache{}
	ac.cacheCollection.Find(bson.M{"api": apiKey, "key": key}).One(&cache)
	if cache.Id == "" {
		response.WriteHeader(404)
		return
	}

	json, _ := json.Marshal(cache)
	response.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(response, "%s", json)
}

func (ac CacheController) Create(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cacheRequest := CacheRequest{}
	json.NewDecoder(request.Body).Decode(&cacheRequest)
	if cacheRequest.CacheValue == "" {
		response.WriteHeader(400)
		return
	}

	apiKey := params.ByName("apikey")
	api := model.Api{}
	apiError := ac.apiCollection.FindId(apiKey).One(&api)
	if apiError != nil {
		response.WriteHeader(404)
		return
	}

	cache := model.Cache{}
	key := params.ByName("key")
	ac.cacheCollection.Find(bson.M{"api": apiKey, "key": key}).One(&cache)
	if cache.Id != "" {
		response.WriteHeader(409)
		return
	}

	cache = model.Cache{
		Id:    bson.NewObjectId(),
		Api:   apiKey,
		Key:   key,
		Value: cacheRequest.CacheValue,
	}
	ac.cacheCollection.Insert(cache)

	response.WriteHeader(200)
}

func (ac CacheController) Update(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cacheRequest := CacheRequest{}
	json.NewDecoder(request.Body).Decode(&cacheRequest)
	if cacheRequest.CacheValue == "" {
		response.WriteHeader(400)
		return
	}

	apiKey := params.ByName("apikey")
	api := model.Api{}
	apiError := ac.apiCollection.FindId(apiKey).One(&api)
	if apiError != nil {
		response.WriteHeader(404)
		return
	}

	cache := model.Cache{}
	key := params.ByName("key")
	cacheError := ac.cacheCollection.Find(bson.M{"api": apiKey, "key": key}).One(&cache)
	if cache.Id == "" {
		fmt.Println("Find cache error:", cacheError)
		response.WriteHeader(404)
		return
	}

	ac.cacheCollection.UpdateId(cache.Id, bson.M{"$set": bson.M{"value": cacheRequest.CacheValue}})

	response.WriteHeader(200)
}

func (ac CacheController) Delete(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	apiKey := params.ByName("apikey")
	api := model.Api{}
	apiError := ac.apiCollection.FindId(apiKey).One(&api)
	if apiError != nil {
		response.WriteHeader(404)
		return
	}

	cache := model.Cache{}
	key := params.ByName("key")
	cacheError := ac.cacheCollection.Find(bson.M{"api": apiKey, "key": key}).One(&cache)
	if cache.Id == "" {
		fmt.Println("Find cache error:", cacheError)
		response.WriteHeader(404)
		return
	}

	ac.cacheCollection.Remove(bson.M{"api": apiKey, "key": key})
	response.WriteHeader(200)
}