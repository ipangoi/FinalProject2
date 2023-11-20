package entity

import (
	"finalProject2/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;type:varchar(191)" json:"username" valid:"required~Your username is required"`
	Email       string        `gorm:"not null;type:varchar(191)" json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password    string        `gorm:"not null" json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age         int           `gorm:"not null;type:int" json:"age"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialmedia"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}
