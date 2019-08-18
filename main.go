package main

import (
	"demo/models"
	_"demo/models"
	_ "demo/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
)

func main() {
	InitSession()
	beego.Run()
}

func InitSession(){
	gob.Register(new(models.User))
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "demo"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./data"
}

