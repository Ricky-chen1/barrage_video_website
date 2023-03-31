package cache

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
)

var Ctx = context.Background()

var (
	RedisDBName    string
	RedisAddr      string
	RedisPassword  string
	RedisNewClient *redis.Client
)

func init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	loadRedis(file)
	redisInit()
}

func loadRedis(file *ini.File) {
	RedisDBName = file.Section("Redis").Key("RedisDBName").String()
	RedisAddr = file.Section("Redis").Key("RedisAddr").String()
	RedisPassword = file.Section("Redis").Key("RedisPassword").String()
}

func redisInit() {
	db, _ := strconv.ParseUint(RedisDBName, 10, 64)
	newClient := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword,
		DB:       int(db),
	})

	if _, err := newClient.Ping(Ctx).Result(); err != nil {
		panic(err)
	}

	RedisNewClient = newClient
}
