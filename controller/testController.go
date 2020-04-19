package controller

import (
	"Wisdomlin/goweb/models"
	"fmt"
	"net/http"
)

type TestController struct {
}

func (t *TestController) Test(w http.ResponseWriter, r *http.Request) {

	testModel := models.TestModel{}
	testModel.Test()

	fmt.Fprint(w, "this is test.test")

}
