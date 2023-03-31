package util

import (
	"barrage_video_website/cache"
	"barrage_video_website/conf"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

var jwtsecret = []byte(conf.JWTSecret)

// GenerateToken 签发用户Token
func GenerateToken(id uint, username string, email string) (string, error) {
	expireTime := conf.TokenExpiredTime
	claims := Claims{
		Id:       id,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  GetUnix(),
			ExpiresAt: GetUnix() + expireTime,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtsecret)
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			return claims, nil
		}
	}
	return nil, err
}

// 检验是否需要重新登录生成token
func CheckIsTokenRefresh(claims *Claims) bool {
	modifyTime := cache.RedisNewClient.Get(cache.Ctx, cache.TokenKey(claims.Id)).Val()
	nowTime, _ := strconv.ParseInt(modifyTime, 10, 64)
	return claims.IssuedAt < nowTime
}
