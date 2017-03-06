package controllers
import (
        . "chatApp/models"
        "strconv"
	"fmt"
)

type GetMsgController struct {
        BaseController
}

type SendMsgController struct {
        BaseController
}


func (this *SendMsgController) Post () {
	uid := this.GetString("uid")
	fid := this.GetString("fid")
	fmt.Println(uid, fid)
	uid64, err := strconv.ParseInt(uid, 10, 64)
	fid64, err := strconv.ParseInt(fid, 10, 64)
	cur_time := this.GetString("cur_time");
	msg := this.GetString("msg");
	id,err := AddMessage(uid64, fid64, msg, cur_time);	
	fmt.Println(id,err)
	this.Data["json"] = map[string]interface{}{"retcode": 0, "message": "ok"}
        this.ServeJSON()
} 

func (this *GetMsgController) Post() {
	uid := this.GetString("uid")
	fid := this.GetString("fid")
	uid64, err := strconv.ParseInt(uid, 10, 64)
	fid64, err := strconv.ParseInt(fid, 10, 64)
	
	msgs,err := GetMessage(uid64, fid64)
	fmt.Println(err)
	this.Data["json"] = map[string]interface{}{"retcode": 0, "message": msgs}
        this.ServeJSON()
        return
}
