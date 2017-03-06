package controllers
import (
	. "chatApp/models"
	"fmt"
)

type LoginUserController struct {
	BaseController
}

type LogoutUserController struct {
	BaseController
}

type RegisterUserController struct {
	BaseController
}

type MyChatController struct {
	BaseController
}


func (this *MyChatController) Get() {
 	check := this.isLogin
        if check {
                str_id := this.GetSession("userLogin")
                user, err := GetUser(str_id.(int64))
                if err == nil {
                        friends, err :=GetAllFriend(str_id.(int64))
                        if err != nil {
                                this.Redirect("/login", 302)
                                return
                        }       
                        this.Data["friends"] = friends
                        this.Data["user"] = user
                        this.TplName = "main.tpl"
                        return	
		}
	} else {
		this.Redirect("/login", 301)
	}
}


func (this *LoginUserController) Get() {
	check := this.isLogin
	if check {
		this.Redirect("/mychat", 302);
		return
	} 
	this.TplName = "login.tpl"
}

func (this *LoginUserController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	if "" == username {
		this.Data["json"] = map[string]interface{}{"code": 0, "message":"please fill the username"} 
		this.ServeJSON()
		return
	}

	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "please fill the passwd"}
		this.ServeJSON()
		return
	}

	err, user := LoginUser(username, password)
	fmt.Println(user)
	if err == nil {
		this.SetSession("userLogin", user[0].Id)
		this.Redirect("/mychat", 302);
		return
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Login failed, username or password is not right"}
	}

	this.ServeJSON()
}

func (this *LogoutUserController) Get() {
	fmt.Println("hello")
	this.DelSession("userLogin");
	this.Redirect("/", 302);	
}

func (this *RegisterUserController) Get() {
	this.TplName = "register.tpl"
}

func (this *RegisterUserController) Post() {
	var user  User
	user.Username = this.GetString("username");
	user.Password = this.GetString("password");
	user.Email    = this.GetString("email");

	if "" == user.Username {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "username is empty"}
		this.ServeJSON()
		return
	}
	
	users, err := GetUserByName(user.Username)
	if len(users) > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "sorory, username: " + user.Username + " is existed"}
                this.ServeJSON()
                return
	}
	
	if "" == user.Password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "password is empty"}
		this.ServeJSON()
		return
	}

	if "" == user.Email {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "Email is empty"}
		this.ServeJSON()
		return
	}

	id, err := AddUser(user)
	if err == nil {
		this.SetSession("userLogin", id)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "add user success", "id": id}
 		this.Redirect("/mychat", 302)
		
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "add user failed"}
		this.ServeJSON()
	}

}

