package controllers

import "github.com/astaxie/beego"

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type APIController struct {
	beego.Controller
}
