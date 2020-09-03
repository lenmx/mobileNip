package routers

import (
	"github.com/astaxie/beego"
	"mobileNip/controllers"
)

func init() {
	initAPI()
}

func initAPI() {
	ns := beego.NewNamespace("/api",
		beego.NSInclude(
			&controllers.APIController{},
		),
	)

	beego.AddNamespace(ns)

	beego.Router("/api/test", &controllers.APIController{}, "GET:Test")
	beego.Router("/api/mobile", &controllers.APIController{}, "GET:GetMobileAttribution")
	beego.Router("/api/ip", &controllers.APIController{}, "GET:GetIpAttribution")
}
