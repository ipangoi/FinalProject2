package entity

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" valid:"required~Message is required" binding:"required"`
	PhotoID uint
	UserID  uint
	User    User
	Photo   Photo
}
