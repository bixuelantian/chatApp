package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "chat.me"
	c.Data["Email"] = "zhangjl613@163.com"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloWorld() {
    main.Data["Website"] = "My Website"
    main.Data["Email"] = "xxx@qq.com"
    main.Data["EmailName"] = "Your Name"
    main.TplName = "main.tpl"
}
