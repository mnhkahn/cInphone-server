package models

type User struct {
	UserName string `json:"username" xorm:"username"`
	Role     string `json:"role" xorm:"role"`
}

type Process struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	Duration int    `json:"duration"`
}

type Record struct {
	Id        int    `json:"id"`
	UserName  string `json:"username" xorm:"username"`
	ProcessId int    `json:"processid" xorm:"processid"`
	State     string `json:"state" xorm:"state"`
}
