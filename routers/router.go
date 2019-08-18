package routers

import (
	"demo/controllers/api"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
   beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
   beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)


   beego.Router("/api/v1/",&api.V1Controller{})
   beego.Router("/api/v1/*",&api.V1Controller{})
}

func TransparentStatic(c *context.Context){
	path := c.Request.URL.Path // /login.html => /static/login.html

	if strings.Index(path,"api") > 0 {
		return
	}

	http.ServeFile(c.ResponseWriter,c.Request,"static"+path)
}



