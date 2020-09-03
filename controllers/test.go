package controllers

func (c *APIController) Test() {
	var res Response
	res = Response{
		Status: "ok",
		Msg:    "success",
		Data:   "hello world.",
	}

	c.Data["json"] = res
	c.ServeJSON()
}
