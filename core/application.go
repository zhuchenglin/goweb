package core

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Application struct {
	Routers     map[string]string
	Controllers map[string]interface{}
}

func (app *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	controlAndFunc := app.Routers[r.URL.Path]
	if controlAndFunc == "" {
		log.Println("路由：" + r.URL.Path + "未注册，请注册后使用！")
		return
	}
	temp := strings.Split(controlAndFunc, "@")
	requestMode, controlString, methodString := strings.ToUpper(temp[0]), temp[1], temp[2]
	if requestMode != r.Method {
		log.Println("路由：" + r.URL.Path + "请求方式不对")
		return
	}
	if controlString == "" || methodString == "" {
		log.Println("路由：" + r.URL.Path + "注册控制器或方法不能为空")
		return
	}
	control := app.Controllers[controlString]
	var result chan interface{} = make(chan interface{})
	go func() {
		_, err := CallFuncByName(control, methodString, w, r)
		result <- err
	}()
	err := <-result
	close(result)
	if err != nil {
		log.Println("路由：" + r.URL.Path + "执行有误")
		log.Println(err)
		return
	}

}

// 反射调用方法
func CallFuncByName(myClass interface{}, funcName string, params ...interface{}) (out []reflect.Value, err error) {
	myClassValue := reflect.ValueOf(myClass)
	m := myClassValue.MethodByName(funcName)
	if !m.IsValid() {
		return make([]reflect.Value, 0), fmt.Errorf("Method not found \"%s\"", funcName)
	}
	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}
	out = m.Call(in)
	return out, nil
}
