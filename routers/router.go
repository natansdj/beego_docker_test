package routers

import (
	"github.com/natansdj/beego_docker/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
