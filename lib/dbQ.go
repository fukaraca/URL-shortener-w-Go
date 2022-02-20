package lib

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//checks whether path is already exist in the DB or not
func checkIfExist(ctx context.Context, path string) bool {
	projection := bson.M{"_id": 0}
	opts := options.FindOne().SetProjection(projection)
	filter := bson.D{{"UrlPath", path}}
	res := Coll.FindOne(ctx, filter, opts)
	resM := bson.M{}
	err := res.Decode(&resM)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Println("decode failed at checkifexist", err)
		ctx.Done()
		return false
	}
	if resM["UrlPath"] == path {
		return true
	}

	return false
}

//inserts new shortened URL to DB and returns error if any
func insertNewURL(ctx context.Context, resp *Response) error {
	document := bson.M{
		"expireAt": resp.ExpireAt,
		"url":      resp.Url,
		"UrlPath":  resp.UrlPath,
		"shortURL": resp.ShortURL,
	}
	_, err := Coll.InsertOne(ctx, document)
	if err != nil {
		log.Println("document couldn't be inserted:", err)
		return err
	}
	return nil
}

//brings url
func getByUrlPath(ctx context.Context, path string) (bool, *Response) {
	projection := bson.M{"_id": 0}
	opts := options.FindOne().SetProjection(projection)
	filter := bson.D{{"UrlPath", path}}
	res := Coll.FindOne(ctx, filter, opts)
	ret := Response{}
	err := res.Decode(&ret)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Println("decode failed at getByUrl:", err)
		ctx.Done()
		return false, nil
	}
	if ret.UrlPath != path {
		return false, nil
	}

	return true, &ret
}
