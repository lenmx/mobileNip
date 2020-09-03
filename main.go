package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "mobileNip/routers"
	"mobileNip/service"
	"mobileNip/util"
	_ "net/http/pprof"
)

func main() {
	util.InitLog()
	service.InitIpResource()
	service.InitMobileResource()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
	}))

	port := beego.AppConfig.String("httpport")
	beego.Run("0.0.0.0:" + port)
}
