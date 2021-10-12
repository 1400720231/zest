package auth

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"math"
	"zset/models/auth"
	"zset/utils"
)

type AuthController struct {
	beego.Controller
}

func (a *AuthController) List() {
	o := orm.NewOrm()

	qs := o.QueryTable("sys_auth")

	auths := []auth.Auth{}
	// 每页显示的条数
	pagePerNum := 2
	// 当前页
	currentPage, err := a.GetInt("page")
	//偏移量
	offsetNum := pagePerNum * (currentPage - 1)
	//搜索关键词
	kw := a.GetString("kw")
	var count int64 = 0

	ret := fmt.Sprintf("当前页;%d,查询条件：%s", currentPage, kw)
	fmt.Println(ret)
	//有查询关键词
	if kw != "" { // 有查询条件的
		// select count(*) from sya_auth where is_delete=0 and username like "%kw"
		count, _ = qs.Filter("is_delete", 0).Filter("auth_name__contains", kw).Count()
		// select count(*) from sya_auth where is_delete=0 and username like "%kw" limit xxx offset xxx
		// 返回对应的结果集对象 相当于django的queryset，当然也可以指定返回值.All(&users, "Id", "Username")
		qs.Filter("is_delete", 0).Filter("auth_name__contains", kw).Limit(pagePerNum).Offset(offsetNum).All(&auths)
	} else { //没有查询关键词
		count, _ = qs.Filter("is_delete", 0).Count()
		qs.Filter("is_delete", 0).Limit(pagePerNum).Offset(offsetNum).All(&auths)

	}
	if err != nil { // 说明没有获取到当前页
		currentPage = 1
	}

	// 总页数=总数/每页个数
	countPage := int(math.Ceil(float64(count) / float64(pagePerNum)))
	//前一页页码
	prePage := 1
	if currentPage == 1 {
		prePage = currentPage
	} else if currentPage > 1 {
		prePage = currentPage - 1
	}
	//下一页页码
	nextPage := 1
	if currentPage < countPage {
		nextPage = currentPage + 1
	} else if currentPage >= countPage {
		nextPage = currentPage
	}
	//分页封装
	page_map := utils.Paginator(currentPage, pagePerNum, count)

	a.Data["auths"] = auths             //序列化之后的查询集合对象 类似django的queryset
	a.Data["prePage"] = prePage         //前一页
	a.Data["nextPage"] = nextPage       //下一页
	a.Data["currentPage"] = currentPage //当前页面
	a.Data["countPage"] = countPage     //总页数
	a.Data["count"] = count             //当前查询条件下的总数
	a.Data["page_map"] = page_map       //封装分页栏的功能 下一页 尾页 首页等功能，相当于django的pure-pagination
	a.Data["kw"] = kw                   //搜索关键词回传显示
	//模板渲染 相当于django的render
	a.TplName = "auth/auth-list.html"

}
