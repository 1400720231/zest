package auth

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

func (a *AuthController) List() {
	a.TplName = "auth/auth-list.html"
}
