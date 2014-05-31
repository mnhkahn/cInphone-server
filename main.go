package main

import (
	"cInphone-server/models"
	_ "cInphone-server/routers"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
	user := models.User{}
	json.Unmarshal(ctx.Input.RequestBody, &user)
	if !(user.Role == "Command" || user.Role == "Coordinate" || user.Role == "Default") {
		ctx.Redirect(500, "")
	}
}

func main() {
	beego.AddFilter("/user", "BeforeExec", FilterUser)
	beego.Run()
}
