package main

import (
	_ "chatApp/routers"
	"github.com/astaxie/beego"
	_ "chatApp/initial"
)

func main() {
	beego.Run()
}

