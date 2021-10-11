package login

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"zset/models/auth"
	"zset/utils"
)

//登陆结构体
type LoginController struct {
	//匿名继承
	beego.Controller
}

//结构体函数封装，Get()方法来自beego.Controller
func (c *LoginController) Get() {

	user_id := c.Ctx.Input.Session("id")
	logs.Error(user_id)
	//>0表示user_id取的到，表示当前有user信息在session中 即是登陆的
	//登陆就直接跳转到首页
	//这是基于接口粒度的控制函数，main.go中可以们用filter做了未登陆状态下的跳转，相当于django的LoginRequiredMixin或者@login_required装饰器
	if user_id != nil {
		c.Redirect(beego.URLFor("HomeController.Get"), 302)
	}
	c.TplName = "login/login.html"
}

//结构体函数封装，Post()方法来自beego.Controller
func (self *LoginController) Post() {
	//获取请求参数
	username := self.GetString("username")
	password := self.GetString("password")
	fmt.Println(username, password)
	//计算密码的md5值
	md5_pwd := utils.GetMd5Str(password)
	fmt.Println(md5_pwd)
	//实例化User model
	userinfo := auth.User{}
	fmt.Println("userinfo", userinfo)
	//声明orm对象
	o := orm.NewOrm()
	//查表 这个写法像flask的orm的写法
	//https://beego.me/docs/mvc/model/query.md#all beego orm文档，毕竟中文文档，这他妈都不好好看还在等什么？
	is_exist := o.QueryTable("sys_user").Filter("user_name", username).Filter("password", md5_pwd).Exist()
	fmt.Println("is_exist", is_exist)

	err := o.QueryTable("sys_user").Filter("user_name", username).Filter("password", md5_pwd).One(&userinfo)
	if err != nil {
		fmt.Println("QueryTable", err)
	}
	fmt.Println("useringo", userinfo)
	//map的value是空接口(interface{})类型，是为了兼容返回体各种格式
	response := map[string]interface{}{} //make(map[string]interface{})

	if !is_exist {
		response["code"] = 600
		response["msg"] = "用户名或密码错误"

	} else if userinfo.IsActive == 0 {
		response["code"] = 600
		response["msg"] = "该用户已停用，请联系管理员"

	} else {
		//把当前登陆的userid写到session 以便下次过来读取
		err := self.SetSession("id", userinfo.Id)

		if err == nil {
			response["code"] = 200
			response["msg"] = "登录成功"

		} else {
			response["code"] = 601
			response["msg"] = "登陆失败"
		}
	}

	//ServeJSON会把self.Data["json"]对应的值以json格式返回，也就是self.Data["json"]的"json"是在框架内
	//强制解析的,是固定值
	self.Data["json"] = response
	self.ServeJSON()

}
