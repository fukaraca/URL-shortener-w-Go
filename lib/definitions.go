package lib

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var Coll *mongo.Collection
var TIMEOUT = 5 * time.Second
var DOMAIN = "localhost:8090"

type request struct {
	url      string
	custom   string
	validFor time.Duration //it is client's responsibility to convert days to hours
}

type Response struct {
	Url      string    `json:"url" bson:"url"`
	UrlPath  string    `json:"UrlPath" bson:"UrlPath"`
	ShortURL string    `json:"shortURL" bson:"shortURL"`
	ExpireAt time.Time `json:"expireAt" bson:"expireAt"`
}
