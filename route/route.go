package route

import "goweb/controller"

var routers = map[string]string{
	// 示例  请求类型@控制器@方法
	"/test":   "get@TestController@Test",
	"/testee": "get@TestController@Testyy",
}

var controllers = map[string]interface{}{
	// 示例  反射中用到
	"TestController": &controller.TestController{},
}

func GetRouter() map[string]string {
	return routers
}

func GetControllers() map[string]interface{} {
	return controllers
}
