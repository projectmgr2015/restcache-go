package controllers

import (
	"../model"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

type ApiController struct {
	apiCollection *mgo.Collection
}

func NewApiController(apiCollection *mgo.Collection) *ApiController {
	return &ApiController{apiCollection}
}

func (ac ApiController) GetApi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uuid := uuid.New()
	api := model.Api{
		Key: uuid,
	}
	ac.apiCollection.Insert(api)
	json, _ := json.Marshal(api)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json)
}
