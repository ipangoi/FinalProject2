package entity

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption  string `gorm:"not null" json:"caption" valid:"required~Description is required"`
	PhotoURL string `gorm:"not null" json:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint
	User     User
}
