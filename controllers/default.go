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

	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
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

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = user.UserName
	this.TplNames = "index.tpl"
}

func (this *MainController) AddProcess() {
	process := models.Process{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &process)

	_, err := engine.Table("process").Insert(&process)

	this.Data["Website"] = process.Duration
	this.Data["Email"] = err
	this.TplNames = "index.tpl"
}

func (this *MainController) RemoveProcess() {
	ProcessId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// process := models.Process{}
	// process.Id = ProcessId
	sql := "delete from process where id=?"
	_, err := engine.Exec(sql, ProcessId)

	this.Data["Website"] = err
	this.Data["Email"] = ProcessId
	this.TplNames = "index.tpl"
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
	id, err := engine.Table("record").Insert(&record)
	this.Data["Website"] = id
	this.Data["Email"] = err
	this.TplNames = "index.tpl"
}

func (this *MainController) GetState() {
	RecordId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	var record = models.Record{Id: RecordId}
	engine.Get(&record)

	this.Data["json"] = record
	this.ServeJson()
}
