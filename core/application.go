package core

import (
	"net/http"
)

type Application struct {
	Routers map[string] interface{}
}

func (app *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//control := app.Routers[r.URL.Path]
	//var baseControl interface{} = &controller.BaseController{}
	//baseControlReflect := reflect.ValueOf(baseControl)
	//m := baseControlReflect.MethodByName("Test")

}
