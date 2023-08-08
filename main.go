package main

import (
	"blog/models"
	_ "blog/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
