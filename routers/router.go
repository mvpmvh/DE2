package routers

import (
	"github.com/mvpmvh/beego"
	"source.discoveryeducation.com/users/mhatch/repos/de2/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.SessionController{}, "get:Login")
	beego.Router("/login", &controllers.SessionController{}, "post:Authenticate")
}
