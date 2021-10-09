package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"zset/controllers/login"
)

func init() {
	beego.Router("/login", &login.LoginController{})
}
