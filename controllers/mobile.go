package controllers

import (
	"mobileNip/service"
	"mobileNip/util"
	"strings"
)

func (c *APIController) GetMobileAttribution() {
	var res Response
	mobile := c.Input().Get("mobile")
	if util.StrIsEmpty(mobile) || len(strings.TrimSpace(mobile)) < 11 {
		res = Response{
			Status: "fail",
			Msg:    "mobile 长度不能小于11位",
			Data:   nil,
		}
	} else {
		mobileItem := service.FindMobileInfoByBinarySearch(mobile, nil)
		res = Response{
			Status: "ok",
			Msg:    "",
			Data:   mobileItem,
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}
