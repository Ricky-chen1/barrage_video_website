package serializer

import "barrage_video_website/model"

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

//带token的用户数据返回
type TokenData struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// 序列化用户
func BuildUser(user *model.User) *User {
	return &User{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
