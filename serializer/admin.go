package serializer

import "barrage_video_website/model"

type Admin struct {
	AdminId uint   `json:"admin_id"`
	Email   string `json:"email"`
}

type AdminTokenData struct {
	Token string `json:"token"`
	Admin Admin  `json:"admin"`
}

func BuildAdmin(admin *model.Admin) *Admin {
	return &Admin{
		AdminId: admin.ID,
		Email:   admin.Email,
	}
}
