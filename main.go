package main

import (
	"log"
	"myapp/db"
	"myapp/routers"
)

func main() {
	db.Init()
	e := routers.Init()

	if err := e.Start(":1234"); err != nil {
		log.Fatal(err)
	}
	//test
}
