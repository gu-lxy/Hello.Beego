package controllers

import (
	"HelloBeego/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
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
func (c *MainController) post()  {
	dataByes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		 c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	var person models.Person
	err = json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
	}
	fmt.Println("用户名：",person.User)
	c.Ctx.WriteString("用户名是："+person.User)

}
