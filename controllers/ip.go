package controllers

import (
	"mobileNip/service"
	"mobileNip/util"
)

func (c *APIController) GetIpAttribution() {
	var res Response
	ip := c.Input().Get("ip")
	if util.StrIsEmpty(ip) {
		res = Response{
			Status: "fail",
			Msg:    "ip 不能为空",
			Data:   nil,
		}
	} else {
		ipItem := service.FindIpInfo(ip)
		res = Response{
			Status: "ok",
			Msg:    "",
			Data:   ipItem,
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}
