package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"zset/controllers"
	"zset/controllers/login"
)

func init() {
	beego.Router("/login", &login.LoginController{})
	beego.Router("/index", &controllers.HomeController{})
}
