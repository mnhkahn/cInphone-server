package models

type User struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
}

type Process struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	Duration string `json:"duration"`
}

type Record struct {
	Id        int    `json:"id"`
	UserName  string `json:"username"`
	ProcessId int    `json:"processid"`
	State     string `json:"state"`
}
