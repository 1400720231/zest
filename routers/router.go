package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"zset/controllers/auth"
	"zset/controllers/index"
	"zset/controllers/login"
	"zset/controllers/logout"
	"zset/controllers/user"
)

func init() {
	beego.Router("/login", &login.LoginController{})
	beego.Router("/logout", &logout.LogoutController{})
	beego.Router("/index", &index.HomeController{})
	// user模块 "get:List"表示http method映射到beego的Controller对应的方法，类似drf的get:list post:create本质上还是http method
	beego.Router("/main/user/list", &user.UserController{}, "get:List")
	beego.Router("/main/user/to_add", &user.UserController{}, "get:ToAdd")
	beego.Router("/main/user/do_add", &user.UserController{}, "post:DoAdd")
	beego.Router("/main/user/is_active", &user.UserController{}, "post:IsActive")
	beego.Router("/main/user/delete", &user.UserController{}, "get:Delete")
	beego.Router("/main/user/reset_pwd", &user.UserController{}, "get:ResetPassword")
	beego.Router("/main/user/to_edit", &user.UserController{}, "get:ToUpdate")
	beego.Router("/main/user/do_edit", &user.UserController{}, "post:DoUpdate")
	beego.Router("/main/user/muli_delete", &user.UserController{}, "post:MuliDelete")
	//auth模块

	beego.Router("/main/user/auth", &auth.AuthController{}, "get:List")

}
