package controllers

import (
	"cInphone-server/models"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
	"net/url"
	"os"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:selinai5@/opensips?charset=utf8")
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

var apikey = "vkl8Pc6QCUjHSemG0wUwVAzQ"
var seckey = "GbrgjnajK6LKl5oa5EIGCuGBRYe8twRP"
var method = "POST"
var url_base = "http://channel.api.duapp.com/rest/2.0/channel/channel"
var url_base1 = "channel.api.duapp.com/rest/2.0/channel/channel"
var query = map[string]string{}

func (this *MainController) Push() {
	title := this.GetString("title")
	description := this.GetString("description")

	query["apikey"] = apikey
	query["message_type"] = "1"
	query["messages"] = "{\"title\":\"" + title + "\",\"description\":\"" + description + "\"}"
	query["method"] = "push_msg"
	query["msg_keys"] = "msgkey"
	query["push_type"] = "3"
	query["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)

	sign := method + url_base
	for k, v := range query {
		sign += k + "=" + v
	}

	sign += seckey

	sign = url.QueryEscape(sign)

	m := md5.New()
	m.Write([]byte(sign))
	sign = hex.EncodeToString(m.Sum(nil))

	url_sign := url_base1 + "?"
	for k, v := range query {
		url_sign += k + "=" + v + "&"
	}

	url_sign += "&sign=" + sign

	req := httplib.Post(url_sign)
	resp, _ := req.Response()

	this.Ctx.ResponseWriter.WriteHeader(resp.StatusCode)
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
