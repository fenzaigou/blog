package models

import (
	utils "blog/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {

	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbuser,
		dbpassword,
		dbhost,
		dbport,
		dbname,
	)

	utils.Debugger("dsn", dsn)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(User), new(Config), new(Post), new(Category), new(Comment))
	orm.RegisterDataBase("default", "mysql", dsn)

	// 自动建表（if not exists）
	orm.RunSyncdb("default", false, true)
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
