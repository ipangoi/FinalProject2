package entity

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" valid:"required~Name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         uint
	User           User
}
