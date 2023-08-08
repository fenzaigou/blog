package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BlogController{}, "*:Home")
	beego.Router("/home", &controllers.BlogController{}, "*:Home")
	beego.Router("/article/:id", &controllers.BlogController{}, "*:Article")
	beego.Router("/article/create", &controllers.BlogController{}, "*:Create")
	beego.Router("/blog/save", &controllers.BlogController{}, "*:Save")
	beego.Router("/blog/update/:id", &controllers.BlogController{}, "*:Update")
	beego.Router("/blog/edit/:id", &controllers.BlogController{}, "*:Edit")
	beego.Router("/blog/delete/:id", &controllers.BlogController{}, "*:Delete")
	beego.Router("/comment/save", &controllers.CommentController{}, "*:CommentSave")
	beego.Router("/comment/delete", &controllers.CommentController{}, "*:CommentDelete")

	// beego.AutoRouter(&controllers.BlogController{})
}
