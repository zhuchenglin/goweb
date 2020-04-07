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
	controllers := route.GetControllers()
	app := new(core.Application)
	app.Routers = routers
	app.Controllers = controllers
	err := http.ListenAndServe("localhost:8080", app)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
