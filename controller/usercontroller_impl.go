package controller

import (
	"finalProject2/database"
	"finalProject2/entity"
	"finalProject2/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct{}

func NewUserHandlerImpl() UserHandler {
	return &UserHandlerImpl{}
}

var (
	appJSON = "application/json"
)

func (s *UserHandlerImpl) UserRegister(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	user := entity.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}
	err := db.Debug().Create(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"age":      user.Age,
	})
}

func (s *UserHandlerImpl) UserLogin(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	password := ""

	user := entity.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password = user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (s *UserHandlerImpl) UserUpdate(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	user := entity.User{}

	userID, _ := strconv.Atoi(c.Param("userID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	user.ID = uint(userID)

	err := db.Model(&user).Where("id = ?", userID).Updates(
		entity.User{
			Username: user.Username,
			Email:    user.Email,
			Age:      user.Age}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
	"id":       user.ID,
	"username": user.Username,
	"email":    user.Email,
	"age":      user.Age,
})
} 
	

func (s *UserHandlerImpl) UserDelete(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	user := entity.User{}

	userID, _ := strconv.Atoi(c.Param("userID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	user.ID = uint(userID)

	err := db.Model(&user).Where("id = ?", userID).Delete(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
