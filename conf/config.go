package conf

import (
	"barrage_video_website/dao"
	"barrage_video_website/pkg/tasks"
	"strings"

	"gopkg.in/ini.v1"
	"gorm.io/gorm/logger"
)

var (
	Path             string
	AppMode          string
	HttpPort         string
	DB               string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	JWTSecret        string
	AdminJWTSecret   string
	TokenExpiredTime int64
)

// 配置文件读取以及路径传递
func Init() {
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadService(cfg)
	LoadMysql(cfg)
	LoadJWT(cfg)
	Path = strings.Join([]string{DBUser, ":", DBPassword, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	//启动定时任务
	tasks.CronJob()
}

func LoadService(cfg *ini.File) {
	AppMode = cfg.Section("service").Key("AppMode").String()
	HttpPort = cfg.Section("service").Key("HttpPort").String()

	//gorm日志等级配置
	if AppMode == "debug" {
		dao.Newlogger = logger.Default.LogMode(logger.Info)
	} else {
		dao.Newlogger = logger.Default
	}
}

func LoadMysql(cfg *ini.File) {
	DB = cfg.Section("mysql").Key("DB").String()
	DBHost = cfg.Section("mysql").Key("DBHost").String()
	DBPort = cfg.Section("mysql").Key("DBPort").String()
	DBUser = cfg.Section("mysql").Key("DBUser").String()
	DBPassword = cfg.Section("mysql").Key("DBPassword").String()
	DBName = cfg.Section("mysql").Key("DBName").String()
}

func LoadJWT(cfg *ini.File) {
	JWTSecret = cfg.Section("jwt").Key("JWTSecret").String()
	AdminJWTSecret = cfg.Section("jwt").Key("AdminJWTSecret").String()
	TokenExpiredTime, _ = cfg.Section("jwt").Key("TokenExpiredTime").Int64()
}
