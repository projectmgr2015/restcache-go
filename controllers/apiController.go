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
	apiRepository *mgo.Collection
}

func NewApiController(apiRepository *mgo.Collection) *ApiController {
	return &ApiController{apiRepository}
}

func (ac ApiController) GetApi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uuid := uuid.New()
	api := model.Api{
		Key: uuid,
	}
	ac.apiRepository.Insert(api)
	json, _ := json.Marshal(api)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json)
}
