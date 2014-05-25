package routers

import (
	"cInphone-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.MainController{}, "post:Add")
	beego.Router("/remove/:id", &controllers.MainController{}, "delete:Remove")
	beego.Router("/get", &controllers.MainController{}, "get:Get")

	beego.Router("/ok/:id", &controllers.MainController{}, "head:OK")
	beego.Router("/state/:id", &controllers.MainController{}, "get:GetState")
}
