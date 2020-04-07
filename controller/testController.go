package controller

import (
	"fmt"
	"net/http"
)

type TestController struct {
	BaseController
}

func (t *TestController) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is test.test")
}