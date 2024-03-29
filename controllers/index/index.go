package index

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"zset/models/auth"
)

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get() {
	user_id2 := h.Ctx.Input.Session("id")

	if user_id2 == nil {
		h.Redirect(beego.URLFor("LoginController.Get"), 301)
	}
	// 后端首页
	o := orm.NewOrm()
	//从session中获取user_id，django的session中间件帮我们做了这一步封装成了request.user
	user_id := h.GetSession("id")
	// interface --> int,静态语言不知道user_id返回的具体类型，GetSession返回的是interface{}数据，所以需要转成int
	user := auth.User{Id: user_id.(int)}
	//外键拼接 类似django的related_name的效果
	o.LoadRelated(&user, "Role")

	//数组对象保存当前用户的角色（一个用户可能多个角色，所以这里是切片）
	auth_arr := []int{}
	for _, role := range user.Role {
		role_data := auth.Role{Id: role.Id}
		o.LoadRelated(&role_data, "Auth")
		for _, auth_date := range role_data.Auth {
			auth_arr = append(auth_arr, auth_date.Id)
		}

	}

	qs := o.QueryTable("sys_auth")

	auths := []auth.Auth{}
	qs.Filter("pid", 0).Filter("id__in", auth_arr).OrderBy("-weight").All(&auths)
	//"select * from sys_user where id in (1,2,3,1)"

	trees := []auth.Tree{}
	//迭代角色 查询角色对应的权限菜单栏 拼接返回给前端展示
	for _, auth_data := range auths { // 一级菜单

		pid := auth_data.Id // 根据pid获取所有的子解点
		tree_data := auth.Tree{Id: auth_data.Id, AuthName: auth_data.AuthName, UrlFor: auth_data.UrlFor, Weight: auth_data.Weight, Children: []*auth.Tree{}}
		GetChildNode(pid, &tree_data)
		trees = append(trees, tree_data)

	}

	h.Data["notify_count"] = 1
	h.Data["trees"] = trees
	h.Data["user"] = user
	h.TplName = "index.html"

}

func (h *HomeController) Welcome() {
	h.TplName = "welcome.html"
}

// 递归
func GetChildNode(pid int, treenode *auth.Tree) {

	o := orm.NewOrm()

	qs := o.QueryTable("sys_auth")
	auths := []auth.Auth{}
	_, err := qs.Filter("pid", pid).OrderBy("-weight").All(&auths)

	if err != nil {
		return
	}

	// 查询三级及以上的菜单
	for i := 0; i < len(auths); i++ {
		pid := auths[i].Id // 根据pid获取所有的子解点
		tree_data := auth.Tree{Id: auths[i].Id, AuthName: auths[i].AuthName, UrlFor: auths[i].UrlFor, Weight: auths[i].Weight, Children: []*auth.Tree{}}
		treenode.Children = append(treenode.Children, &tree_data)
		GetChildNode(pid, &tree_data)
	}

	return

}
