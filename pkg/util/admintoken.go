package util

import (
	"barrage_video_website/conf"

	"github.com/dgrijalva/jwt-go"
)

type AdminClaims struct {
	AdminId uint
	jwt.StandardClaims
}

var adminjwtsecret = []byte(conf.AdminJWTSecret)

// GenerateToken 签发用户Token
func GenerateAdminToken(admin_id uint) (string, error) {
	expireTime := conf.TokenExpiredTime
	admin_claims := AdminClaims{
		AdminId: admin_id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  GetUnix(),
			ExpiresAt: GetUnix() + expireTime,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, admin_claims)
	token, err := tokenClaims.SignedString(adminjwtsecret)
	return token, err
}

// ParseToken 验证用户token
func ParseAdminToken(token string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return adminjwtsecret, nil
	})
	if tokenClaims != nil {
		if admin_claims, ok := tokenClaims.Claims.(*AdminClaims); ok {
			return admin_claims, nil
		}
	}
	return nil, err
}
