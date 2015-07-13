package model
import "gopkg.in/mgo.v2/bson"

type Cache struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Api   string `json:"api" bson:"api"`
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}
