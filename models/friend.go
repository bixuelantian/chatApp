package models
import (
        "github.com/astaxie/beego/orm"
        _ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Friend struct {
	Id int64
	Userid int64
	Friendid int64
}

type FMsg struct {	
	Id string
	Username string
	Email  string
	Rid   string
	Msgnum string
}

func (this *Friend) TableName() string {
        return "tb_friend"
}

func DelFriend(rid int64) (num int64, err error) {
	o := orm.NewOrm()
	friend := Friend{Id:rid}
	num, err = o.Delete(&friend)
	return num, err
}

func AddFriend(c_us Friend) (int64, error) {
        o := orm.NewOrm()
        o.Using("default")
        id, err := o.Insert(&c_us)
        return id, err
}

func GetFriend(uid int64, fid int64) (friends []Friend, err error){
	o := orm.NewOrm()
        qs := o.QueryTable("tb_friend")
        cond := orm.NewCondition()
        cond = cond.And("userid", uid)
        cond = cond.And("friendid", fid)
        qs = qs.SetCond(cond)
        err = qs.Limit(1).One(&friends)
        return friends, err
}

func GetAllFriend(uid int64) (frieds map[string]FMsg, err error){
	o := orm.NewOrm()
	var maps []orm.Params
	sSql := "select user.*, friend.id as rid from tb_user user, tb_friend friend where friend.friendid = user.id and friend.userid = ?"
	num, err := o.Raw(sSql, uid).Values(&maps)
	fmt.Println(num)
	var fmsg map[string]FMsg
	fmsg = make(map[string]FMsg)
	for index,item := range maps {
		msg := FMsg{item["id"].(string), item["username"].(string), item["email"].(string), item["rid"].(string), "0"}
		fmsg[item["id"].(string)] = msg
		fmt.Println(index)
	}

	sSql = "select sender, count(*) as msgnum from tb_message msg where msg.receiver=? AND msg.status=0 group by sender"
	var msgs []orm.Params
	num, err = o.Raw(sSql, uid).Values(&msgs)
	for index,item := range msgs {
		sender := item["sender"].(string)
		if _,ok := fmsg[sender]; ok {
			msg := fmsg[sender]
			msg.Msgnum = item["msgnum"].(string)
			fmsg[sender] = msg
                	fmt.Println(index)
		}
        }

	fmt.Println(fmsg)
	
	return fmsg, err
}


func init() {
        orm.RegisterModel(new(Friend))
}
