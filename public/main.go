package main

import (
	"log"
	"net/http"
	"warehouse/core"
	"warehouse/route"
)

func main() {
	Start()
}

func Start() {
	routers := route.GetRouter()
	app := new(core.Application)
	app.Routers = routers
	err := http.ListenAndServe("localhost:8080", app)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
