package route

import (
	"wycto/app/controller"
	"wycto/weigo"
)

func init() {
	weigo.Router("/", &controller.IndexController{})
	weigo.Router("/user/", &controller.UserController{})
	weigo.Router("/api/", &controller.APIController{})
	weigo.Router("/test/", &controller.TestController{})
	weigo.Router("/model/", &controller.ModelController{})
}
