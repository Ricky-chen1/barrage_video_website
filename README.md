# barrage_video_website 弹幕视频网站

**此项目使用Gin+Gorm+ ，基于RESTful API实现的一个简单的弹幕视频网站后端**  

## 项目主要模块介绍  
* 1.用户模块  
* 2.视频模块 
* 3.管理员模块

## 项目功能简要介绍
* 1.视频点击量排行榜(redis缓存实现)
* 2.视频搜索功能(mysql模糊搜索实现)

## 项目主要依赖  
* Gin 
* Gorm
* ini  
* Mysql
* redis
* cron  
* jwt-go  

## 项目结构  
```
barrage_video_website/
├── cache
├── conf
├── controllers
├── dao
├── middleware
├── model
├── pkg
│  ├── errno  
│  └── tasks
│  └── util
├── routers  
├── serializer    
├── service
├── tmp
└── upload
```     

* cache : 用于redis缓存
* conf : 配置文件的配置和读入
* controllers : 用于定义接口函数
* dao : mysql数据库交互    
* middleware : 应用中间件    
* model : 应用数据库模型 
* pkg\errno : 错误码和错误信息封装
* pkg\tasks : 任务实现,如定时任务 
* pkg\util : 工具函数实现  
* routers : 路由逻辑处理  
* serializer : 将数据序列化为 json 的函数(序列化器)      
* service : 接口函数的实现(业务逻辑)  
* tmp : 热加载包         

## 项目运行  
**此项目使用Go Mod管理依赖。**    
### 下载依赖    
`go mod tidy`  
### 运行  
`go run main.go`
