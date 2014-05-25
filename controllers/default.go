package controllers

import (
	"cInphone-server/models"
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) AddUser() {
	user := models.User{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = user.UserName
	this.TplNames = "index.tpl"
}

func (this *MainController) AddProcess() {
	process := models.Process{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &process)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = process.Content
	this.TplNames = "index.tpl"
}

func (this *MainController) RemoveProcess() {
	ProcessId := this.Ctx.Input.Param(":id")
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = ProcessId
	this.TplNames = "index.tpl"
}

func (this *MainController) GetProcess() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) OK() {
	ProcessId := this.Ctx.Input.Param(":id")
	UserName := this.Ctx.Input.Param(":user")
	this.Data["Website"] = ProcessId
	this.Data["Email"] = UserName
	this.TplNames = "index.tpl"
}

func (this *MainController) GetState() {
	RecordId := this.Ctx.Input.Param(":id")
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = RecordId
	this.TplNames = "index.tpl"
}
