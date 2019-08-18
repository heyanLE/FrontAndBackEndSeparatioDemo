package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	)

type User struct {
	Id 				int 			`json:"id"`
	Username		string 			`json:"username"`
	PasswordHash	string 			`json:"password_hash"`
}

func init(){
	dbHost := beego.AppConfig.String("dbhost")
	dbPort := beego.AppConfig.String("dbport")
	dbUser := beego.AppConfig.String("dbuser")
	dbPass := beego.AppConfig.String("dbpass")
	dbName := beego.AppConfig.String("dbname")
	dbStr := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/"+dbName + "?charset=utf8"
	e := orm.RegisterDataBase("default","mysql",dbStr)
	if e != nil {
		panic(e.Error())
	}
	orm.RegisterModel(new(User))
	e = orm.RunSyncdb("default",true,true)
	if e != nil {
		panic(e.Error())
	}
}
