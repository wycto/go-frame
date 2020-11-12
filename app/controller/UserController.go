package controller

import (
	"io"
	"wycto/weigo"
)

type UserController struct {
	weigo.Controller
	modelName string
}

func (c UserController) Index() {
	io.WriteString(c.Context.ResponseWriter, "Welcome UserController Index")
}

func (c *UserController) Login() {
	io.WriteString(c.Context.ResponseWriter, "this is user login"+c.Context.Request.Method)
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.FormValue("name"))
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.Form.Get("name"))
	io.WriteString(c.Context.ResponseWriter, "ControllerName:"+c.Context.ControllerName)
	io.WriteString(c.Context.ResponseWriter, "ActionName:"+c.Context.ActionName)
	io.WriteString(c.Context.ResponseWriter, "name:"+c.Context.Request.Form.Get("name"))
}

func (c *UserController) UserInfo() {
	io.WriteString(c.Context.ResponseWriter, "this is userinfo")
}
