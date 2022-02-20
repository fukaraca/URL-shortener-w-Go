package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"io"
	"log"
	"os"
)

func InitServer() *gin.Engine {

	//logger middleware teed to log.file
	logfile, err := os.OpenFile("./logs/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Could not create/open log file")
	}
	errlogfile, err := os.OpenFile("./logs/errlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Could not create/open err log file")
	}
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(errlogfile, os.Stdout)

	//starts with builtin Logger() middleware
	return gin.Default()

}

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://rootuser:rootpass@localhost:27017"))
	if err != nil {
		log.Fatalln("connection failed:", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	Coll = client.Database("URLs").Collection("short-urls")

	//assign unending TTL for every documents by default
	keys := bsonx.Doc{{Key: "expireAt", Value: bsonx.Int32(1)}}
	ttl := int32(0)

	ind := mongo.IndexModel{
		Keys:    keys,
		Options: &options.IndexOptions{ExpireAfterSeconds: &ttl},
	}
	_, err = Coll.Indexes().CreateOne(ctx, ind)
	if err != nil {
		log.Println("Error occurred while creating index", err)
	} else {
		log.Println("Index creation success")
	}

}
