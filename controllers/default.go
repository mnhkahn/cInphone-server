package controllers

import (
	"cInphone-server/models"
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
	"os"
	"strconv"
)

type MainController struct {
	beego.Controller
}

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:qwerty@/opensips?charset=utf8")
	// defer engine.Close()
	engine.ShowSQL = true

	f, err := os.Create("sql.log")
	if err != nil {
		panic(err.Error())
		return
	}
	engine.Logger = f

}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) AddUser() {
	user := models.User{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	_, err := engine.Table("user").Insert(&user)

	if err == nil {
		this.Data["json"] = ""
		this.ServeJson()
	} else {
		this.Abort("500")
	}
}

func (this *MainController) GetUsers() {
	users := []models.User{}
	username := this.Ctx.Input.Param(":user")
	engine.Table("user").Where("username!=?", username).Find(&users)
	this.Data["json"] = users
	this.ServeJson()
}

func (this *MainController) AddProcess() {
	process := models.Process{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &process)

	_, err := engine.Table("process").Insert(&process)

	if err == nil {
		this.Data["json"] = ""
		this.ServeJson()
	} else {
		this.Abort("500")
	}
}

func (this *MainController) RemoveProcess() {
	ProcessId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// process := models.Process{}
	// process.Id = ProcessId
	sql := "delete  from process where id=?"
	_, err := engine.Exec(sql, ProcessId)

	if err == nil {
		this.Data["json"] = ""
		this.ServeJson()
	} else {
		this.Abort("500")
	}
}

func (this *MainController) GetProcess() {
	processes := []models.Process{}
	engine.Table("process").OrderBy("id").Find(&processes)
	this.Data["json"] = processes
	this.ServeJson()
}

func (this *MainController) OK() {
	ProcessId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	UserName := this.Ctx.Input.Param(":user")

	record := models.Record{}
	record.UserName = UserName
	record.ProcessId = ProcessId
	record.State = "ok"
	_, err := engine.Table("record").Insert(&record)
	if err == nil {
		this.Data["json"] = ""
		this.ServeJson()
	} else {
		this.Abort("500")
	}
}

func (this *MainController) GetState() {
	RecordId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	var record = models.Record{Id: RecordId}
	engine.Get(&record)

	this.Data["json"] = record
	this.ServeJson()
}
