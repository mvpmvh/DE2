package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"source.discoveryeducation.com/users/mhatch/repos/de2/factories"
	"source.discoveryeducation.com/users/mhatch/repos/de2/models"
	"source.discoveryeducation.com/users/mhatch/repos/de2/services"
)

type SessionController struct {
	beego.Controller
	userService *services.UserService
}

func (this *SessionController) Prepare() {
	if this.userService == nil {
		this.userService = factories.NewUserService()
	}
}

func (this *SessionController) Authenticate() {
	user := new(models.User)

	if err := this.ParseForm(user); err != nil {
		panic(err)
	}
	fmt.Printf("user %#v\n", user)
	criteria := models.Criteria{"user": user, "referer": this.Ctx.Request.Referer()}

	if user = this.userService.FindUser(criteria); user.Id == nil {
		this.Data["ErrorMessage"] = "Invalid username/password"
	} else {
		this.Data["Message"] = "Success!"
	}

	this.TplNames = "home.tpl"
}

func (this *SessionController) Login() {
	this.TplNames = "home.tpl"
}
