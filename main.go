package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"zset/utils"

	//beego2和之前的用法有些不一样 请尽量参考官方文档
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	//把model对应的包导入不然不会执行init函数 model无法注册被发现 有点像django的app_installed
	_ "zset/models/auth"
	_ "zset/models/caiwu"
	_ "zset/models/my_center"
	_ "zset/models/news"
	_ "zset/routers"
)

func init() {
	//beego2返回两个参数，对参数名找不到的情况做了错误处理
	username, _ := beego.AppConfig.String("username")
	pwd, _ := beego.AppConfig.String("pwd")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	db, _ := beego.AppConfig.String("db")

	// root:250onioN!!!!@tcp(localhost:3306)/zset?charset=utf8
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"

	err1 := orm.RegisterDriver("mysql", orm.DRMySQL)
	//这里的"default"相当于一个别名，万一一个是主一个是从数据库，需要读写分离的时候就可以根据这个做映射
	err2 := orm.RegisterDataBase("default", "mysql", dataSource)
	fmt.Println(dataSource)
	fmt.Println(err1, err2)
	//\n可以让输出的时候没有%出现
	fmt.Printf("host:%s|port:%s|db:%s\n", host, port, db)

}

func main() {
	//开启session 或者配置文件写：sessionon = true
	/*
		如果不开启这个配置但是你用了self.Session.set相关的用法，会报指针错误：invalid memory address or nil pointer dereference
		但是其实不是语法错误只是你没配置开启session，如果你对整个框架熟不熟悉会一直以为是语法错误，很难排查。。。
	*/
	beego.BConfig.WebConfig.Session.SessionOn = true
	// 对index页面进行未登录请求拦截,BeforeRouter的作用个生命周期的意义可以查看beego官网的beego架构图
	beego.InsertFilter("/main*", beego.BeforeRouter, utils.LoginFilter)
	//数据库命令行迁移
	//orm.RunCommand()
	//直接执行数据库迁移操作
	//err := orm.RunSyncdb("default", false, true)
	//fmt.Println(err)
	orm.Debug = true
	//日志配置
	//声明一个日志对象，下面有一个日志处理器，类型为文件类型的日志，文件配置的位置在当前目录的logs文件名为log.out 有两个
	//日志级别error info,在logs下面会有三个文件：log.out log.error.out log.info.out
	//log.out>log.error.out+log.info.out log.out还保存了beego的请求日志
	//beego每次启动的日志会保存在log.info.out中
	//其他的日志信息就是你调用logs.Error logs.Info对应的日志记录内容了
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/log.out","separate":["error","info"]}`)
	beego.Run()
	//res := utils.GetMd5Str("admin123")
	//fmt.Println(res)

}
