package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"zset/controllers/logout"

	"zset/controllers"
	"zset/controllers/login"
)

func init() {
	beego.Router("/login", &login.LoginController{})
	beego.Router("/logout", &logout.LogoutController{})
	beego.Router("/index", &controllers.HomeController{})
}
