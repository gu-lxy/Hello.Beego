package main

import (
	_ "HelloBeego/routers"
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "syscall"
)

func main() {
	config := beego.AppConfig//定义config变量，接受并赋值为全局配置变量
	appName := config.Strings("appname")
	fmt.Println("项目应用名称：",appName)
	port,err := config.Int("httpport")
	if err != nil {
		panic("项目配置错误，请重试")
		return
	}
	fmt.Println("应用监听窗口:",port)
	driver := config.Strings("db_driver")//数据库驱动
	dbUser := config.Strings("db_user")//数据库用户名
	dbPasssword := config.Strings("db_password")//密码
	dbip := config.Strings("db_ip")
	dbName := config.Strings("db_name")

	db, err := sql.Open(driver+":"+dbUser+":" +dbPasssword+"@tcp(" +dbip+")/"+dbName+"?charset=utf8")
	if err != nil {
		panic("数据库连接失败。请重新尝试")
	}
	fmt.Println(db)

	beego.Run()
}

