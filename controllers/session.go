package controllers

import (
	_ "github.com/mvpmvh/beego"
	"source.discoveryeducation.com/users/mhatch/repos/de2/factories"
	"source.discoveryeducation.com/users/mhatch/repos/de2/models"
	"source.discoveryeducation.com/users/mhatch/repos/de2/services"
)

type SessionController struct {
	baseController
	userService *services.UserService
}

func (this *SessionController) Prepare() {
	this.baseController.Prepare()

	if this.userService == nil {
		this.userService = factories.NewUserService()
	}
}

func (this *SessionController) Authenticate() {
	user := new(models.User)

	if err := this.ParseForm(user); err != nil {
		panic(err)
	}
	//fmt.Printf("user %#v\n", user)
	criteria := models.Criteria{"user": user, "referer": this.Ctx.Request.Referer()}

	if user = this.userService.FindUser(criteria); user.Id == nil {
		this.Data["ErrorMessages"] = []string{"Invalid username/password"}
	} else {
		//TODO: redirect & implement jeff's landing page (BED-20)'
	}

	this.TplNames = "public/login.tpl"
}

func (this *SessionController) Login() {
	this.TplNames = "public/login.tpl"
}
