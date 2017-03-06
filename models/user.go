package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id	 int64
	Username string
	Password string
	Email    string
	Created  string
}

func (this *User) TableName() string {
	return "tb_user"
}

func LoginUser(username string, password string) (err error, user []User) {
	o := orm.NewOrm()
	qs := o.QueryTable("tb_user")
	cond := orm.NewCondition()

	cond = cond.And("username", username)
	cond = cond.And("password", password)

	qs = qs.SetCond(cond)
	var users []User
	err1 := qs.Limit(1).One(&users)
	return err1, users
}

func AddUser(c_us User) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	n_us := new(User)

	n_us.Username = c_us.Username
	n_us.Password = c_us.Password
	n_us.Email = c_us.Email
	
	id, err := o.Insert(n_us)
	return id, err
}

func GetUserByName(name string) (user []User, err error) {
	o := orm.NewOrm()
        qs := o.QueryTable("tb_user")
        cond := orm.NewCondition()
        cond = cond.And("username", name)
        qs = qs.SetCond(cond)
        var users []User
        err1 := qs.Limit(1).One(&users)
        return  users, err1
}

func GetUser(id int64) (User, error) {
        var user User
        var err error
        o := orm.NewOrm()
        user = User{Id: id}
        err = o.Read(&user)
        if err == orm.ErrNoRows {
                return user, nil
        }
        return user, err
}

func GetAllUser() (users []User, err error) {
	o := orm.NewOrm()
	sSql := "SELECT * FROM tb_user"
	num, err := o.Raw(sSql).QueryRows(&users)
	if err == nil {
        	fmt.Println("users nums: %d", num)
    	}
	return users, err 
}

func init() {
	orm.RegisterModel(new(User))
}

