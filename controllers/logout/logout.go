package logout

import beego "github.com/beego/beego/v2/server/web"

//登陆结构体
type LogoutController struct {
	//匿名继承
	beego.Controller
}

//request并不是固定值 还是那句话相当于python继承中的self
func (request *LogoutController) Get() {
	//删除session中的user_id
	request.DelSession("id")
	//重定向到登陆页面
	request.Redirect(beego.URLFor("LoginController.Get"), 302)

}
