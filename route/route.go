package route

var routers = map[string] string{
	"/test": "TestController.Test",
}

func GetRouter() map[string]string {
	return routers
}
