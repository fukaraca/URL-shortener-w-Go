package main

import (
	"URL-shortener-w-Go/lib"
	"log"
)

var r = lib.InitServer()

func init() {
	lib.ConnectDB()
}

func main() {
	r.GET("/:shortURL", lib.BringMeLonger)
	r.POST("/squeeze", lib.MakeThisShorter)

	log.Fatalln("router failed:", r.Run(":8090"))

}
