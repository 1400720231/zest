package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"zset/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
