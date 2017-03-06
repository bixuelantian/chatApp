package routers

import (
	"chatApp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.MainController{}, "get:HelloWorld")
    beego.Router("/register", &controllers.RegisterUserController{})
    beego.Router("/login", &controllers.LoginUserController{})
    beego.Router("/logout", &controllers.LogoutUserController{})
    beego.Router("/mychat", &controllers.MyChatController{})
    beego.Router("/addfriend", &controllers.AddFriendController{})
    beego.Router("/delfriend", &controllers.DelFriendController{})
    beego.Router("/getmsg", &controllers.GetMsgController{})
    beego.Router("/sendmsg", &controllers.SendMsgController{})
}
