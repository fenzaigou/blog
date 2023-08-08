package controllers

import (
	"blog/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()

	// 后台相关
	if p.controllerName == "admin" && p.actionName != "login" {
		if p.GetSession("user") == nil {
			p.History("未登录", "/admin/login")
		}
	}

	// 初始化前台页面相关元素
	if p.controllerName == "blog" {
		p.Data["actionName"] = p.actionName
		var result []*models.Config

		// 从数据库里获取配置，存到 controller 的 configs 里
		p.o.QueryTable(new(models.Config).TableName()).All(&result)
		configs := make(map[string]string)
		for _, v := range result {
			configs[v.Name] = v.Value
		}
		p.Data["configs"] = configs
	}
}

func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString(
			fmt.Sprintf(
				`<script>alert('%s');window.history.go(-1);</script>`,
				msg,
			),
		)
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}

func (p *baseController) getClientIp() string {
	fmt.Printf("RemoteAddr in ctx request: %v\n", p.Ctx.Request.RemoteAddr)
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func getIdFromRequest(c *baseController) (int, error) {
	return strconv.Atoi(c.Ctx.Input.Params()[":id"])
}
