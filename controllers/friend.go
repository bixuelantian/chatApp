package controllers
import (
	. "chatApp/models"
	"fmt"
	"strconv"
)

type AddFriendController struct {
	BaseController
}

type DelFriendController struct {
	BaseController
}

func (this *DelFriendController) Post() {
	rid := this.GetString("rid")
	rid64, err := strconv.ParseInt(rid, 10, 64)
	num, err := DelFriend(rid64)
	if err == nil {
		fmt.Println(num)
		this.Data["json"] = map[string]interface{}{"retcode":0, "retmsg":"ok"}
	}else{
		this.Data["json"] = map[string]interface{}{"retcode":1}
	}
        this.ServeJSON()
}

func (this *AddFriendController) Post() {
	uid := this.GetString("userid")
	fname := this.GetString("friendname")
	users, err := GetUserByName(fname)
	fmt.Println(err)
	if len(users) == 0 {
		this.Data["json"] = map[string]interface{}{"retcode":1, "retmsg":"Sorry , Username " + fname + " not existed"}
		this.ServeJSON() 	
		return 
	}
	uid64, err := strconv.ParseInt(uid, 10, 64)
	
	if uid64 == users[0].Id {
		this.Data["json"] = map[string]interface{}{"retcode":1, "retmsg":"Sorry , you cant add your self"}
                this.ServeJSON()
                return
	}

	friends, err := GetFriend(uid64, users[0].Id)
	if len(friends) > 0 {
		 this.Data["json"] = map[string]interface{}{"retcode":1, "retmsg":fname + " has been your friend"}
                this.ServeJSON()
                return

	}
	fmt.Println("world")
	var friend1,friend2 Friend
	friend1.Userid = uid64
	friend1.Friendid = users[0].Id
	id, err := AddFriend(friend1)
	friend2.Userid = users[0].Id
	friend2.Friendid = uid64
	id, err = AddFriend(friend2)

	fmt.Println(id)
	this.Data["json"] = map[string]interface{}{"retcode":0, "retmsg":"ok"}
	this.ServeJSON() 	
}
