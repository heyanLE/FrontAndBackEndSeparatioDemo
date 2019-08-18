package api

import (
	"demo/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

const SessionUserKey  =  "cn.heyanle.demo.SESSION_USER_KEY"

type V1Controller struct {
	beego.Controller
}

type UserMessage struct {
	Username 	string 		`json:"username"`
	Password 	string 		`json:"password"`
}

type ResponseBody struct {
	Code 		int 		`json:"code"`
	Message 	string 		`json:"message"`
	Value 		interface{} `json:"value"`
}

func (c *V1Controller) Post()  {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/api/v1/login":
		c.LoginPost()
		break
	case "/api/v1/register":
		c.RegisterPost()
		break
	}
}

func (c *V1Controller) Get()  {
	path := c.Ctx.Request.URL.Path
	switch path {
	case "/api/v1/user":
		c.UserGet()
		break
	case "/api/v1/login-out":
		c.LoginOutGet()
		break
	}
}

func (c *V1Controller) LoginPost (){
	body := c.Ctx.Input.RequestBody
	um := UserMessage{}
	e := json.Unmarshal(body,&um)
	rb := ResponseBody{}
	if e != nil {
		rb.Code = 400
		rb.Message = "参数错误"
		c.Data["json"] = &rb
		c.ServeJSON()
	}
	u,e := models.Login(um.Username,um.Password)
	if e == models.UsernameNotFind {
		rb.Code = 404
		rb.Message = "用户名不存在"
	} else if e == models.PasswordError{
		rb.Code = 404
		rb.Message = "密码错误"
	} else if e == nil {
		rb.Code = 200
		rb.Message = "登陆成功"
		c.SetSession(SessionUserKey,&u)
	} else {
		rb.Code = 500
		rb.Message = e.Error()
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}

func (c *V1Controller) RegisterPost (){
	body := c.Ctx.Input.RequestBody
	um := UserMessage{}
	e := json.Unmarshal(body,&um)
	rb := ResponseBody{}
	if e != nil {
		rb.Code = 400
		rb.Message = "参数错误"
		c.Data["json"] = &rb
		c.ServeJSON()
	}
	_,e = models.Register(um.Username,um.Password)
	if e == models.UserExist {
		rb.Code = 404
		rb.Message = "用户已存在"
	}else if e == nil {
		rb.Code = 200
		rb.Message = "注册成功"
	}else {
		rb.Code = 500
		rb.Message = e.Error()
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}

func (c *V1Controller) UserGet (){
	u := c.GetSession(SessionUserKey)
	rb := ResponseBody{}
	if u == nil {
		rb.Code = 404
		rb.Message = "当前没有登录用户"
	}else {
		rb.Code = 200
		rb.Message = "获取登录用户成功"
		rb.Value = &u
	}
	c.Data["json"] = &rb
	c.ServeJSON()
}

func (c *V1Controller) LoginOutGet (){
	c.DelSession(SessionUserKey)
	c.Data["json"] = &ResponseBody{Code:200,Message:"登出成功"}
	c.ServeJSON()
}
