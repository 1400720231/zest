package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "zset/routers"
)

func init() {
	username := beego.AppConfig.String("username")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db := beego.AppConfig.String("db")

	// username:pwd@tcp(ip:port)/db?charset=utf8&loc=Local
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8mb4"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//这里的"default"相当于一个别名，万一一个是主一个是从数据库，需要读写分离的时候就可以根据这个做映射
	orm.RegisterDataBase("default", "mysql", dataSource)

	fmt.Sprintf("host:%s|port:%s|db:%s", host, port, db)

}

func main() {
	//数据库命令行迁移
	orm.RunCommand()
	beego.Run()
}
