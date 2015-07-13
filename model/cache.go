package model

type Cache struct {
	Id    string `json:"id" bson:"_id"`
	Api   string `json:"api" bson:"api"`
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}
