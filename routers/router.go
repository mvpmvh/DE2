package routers

import (
	"github.com/mvpmvh/beego"
	"source.discoveryeducation.com/users/mhatch/repos/de2/controllers"
	"source.discoveryeducation.com/users/mhatch/repos/de2/taglib"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.SessionController{}, "get:Login")
	beego.Router("/login", &controllers.SessionController{}, "post:Authenticate")

	beego.AddFuncMap("params", taglib.Params)
}
