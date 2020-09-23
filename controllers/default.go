package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	if name != "gu"|| age !="99" || sex != "name" {
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据准确"))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
/*
*该post的方法是处理post类型的请求的时候要调用的方法
*/
func (c *MainController) Post() {
	fmt.Println("post类型的请求")
	c.Data["Website"] = "beego.me"
	user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为",user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是",psd)
	if user!= "gu"||psd!="123456" {
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据准确"))
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
