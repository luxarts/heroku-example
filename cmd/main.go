package main

import (
	"heroku-example/internal/router"
	"log"
)

func main(){
	r := router.Init()

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
