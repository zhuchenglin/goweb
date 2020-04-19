package main

import (
	"Wisdomlin/goweb/core"
	"Wisdomlin/goweb/route"
	"log"
	"net/http"
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
		log.Println("ListenAndServe: ", err)
	}
}
