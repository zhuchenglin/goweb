package controller

import (
	"fmt"
	"goweb/models"
	"net/http"
)

type TestController struct {
}

func (t *TestController) Test(w http.ResponseWriter, r *http.Request) {

	testModel := models.TestModel{}
	testModel.Test()

	fmt.Fprint(w, "this is test.test")

}
