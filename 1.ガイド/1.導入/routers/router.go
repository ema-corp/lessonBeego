package routers

import (
	"lessonBeego/1.ガイド/1.導入/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
