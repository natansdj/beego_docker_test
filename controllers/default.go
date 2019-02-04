package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vm"
	c.Data["Email"] = "test@example.com"
	c.TplName = "index.tpl"
}
