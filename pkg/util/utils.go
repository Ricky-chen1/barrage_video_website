package util

import (
	"math/rand"
	"time"
)

// 随机字符串生成
func RandStringByte(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	res := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取年月日
func GetDay() string {
	templates := "20230312"
	return time.Now().Format(templates)
}
