package models
 
import (
        "github.com/astaxie/beego/orm"
//        _ "github.com/go-sql-driver/mysql"
        "fmt"
)

type Message struct {
        Id int64
        Sender int64
        Receiver int64
	Message string
	Sendtime string
	Status   int
}

func (this *Message) TableName() string {
        return "tb_message"
}

func AddMessage(uid int64, fid int64, msg string, cur_time string) (id int64, err error){
	o := orm.NewOrm()
        o.Using("default")
        n_m := new(Message)

	n_m.Sender = uid
	n_m.Receiver = fid
	n_m.Message = msg	
	n_m.Sendtime = cur_time;
        id, err = o.Insert(n_m)
        return id, err
}

func GetMessage(uid int64, fid int64 ) (msg []Message, err error) {
	o := orm.NewOrm()
        sSql := "SELECT * FROM tb_message where (sender=? And receiver=?) or (sender=? and receiver=?)"
        num, err := o.Raw(sSql, uid, fid, fid, uid).QueryRows(&msg)
        if err == nil {
                fmt.Println("users nums:", num)
        }
	fmt.Println(msg)

	for i:=0; i<len(msg); i++ {
		mm := msg[i];
		mm.Status = 1;		
		o.Update(&mm)
	}	
        return msg, err
}

func init() {
        orm.RegisterModel(new(Message))
}
