package main

import (
	"barrage_video_website/conf"
	"barrage_video_website/dao"
	"barrage_video_website/routers"
)

func main() {
	//读取配置文件
	conf.Init()
	//连接数据库
	dao.MysqlInit(conf.Path)
	//路由服务
	routers.RoutersInit(conf.HttpPort)
}
