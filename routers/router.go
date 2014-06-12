package routers

import (
	"cInphone-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/push", &controllers.MainController{}, "get:Push")

	beego.Router("/user", &controllers.MainController{}, "post:AddUser")
	beego.Router("/users/:user", &controllers.MainController{}, "get:GetUsers")

	beego.Router("/process", &controllers.MainController{}, "post:AddProcess")
	beego.Router("/process/:id", &controllers.MainController{}, "delete:RemoveProcess")
	beego.Router("/process", &controllers.MainController{}, "get:GetProcess")

	beego.Router("/record/:id/:user", &controllers.MainController{}, "post:OK")
	beego.Router("/record/:id", &controllers.MainController{}, "get:GetState")
}
