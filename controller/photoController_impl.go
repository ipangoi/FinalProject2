package controller

import (
	"finalProject2/database"
	"finalProject2/entity"
	"finalProject2/helper"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoHandlerImpl struct{}

func NewPhotoHandlerImpl() PhotoHandler {
	return &PhotoHandlerImpl{}
}

func (s *PhotoHandlerImpl) PhotoCreate(c *gin.Context) {
	var db = database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	Photo := entity.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func (s *PhotoHandlerImpl) PhotoGet(c *gin.Context) {
	var db = database.GetDB()
	//userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	var Photo []entity.Photo
	//userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email")
	}).Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func (s *PhotoHandlerImpl) PhotoGetOne(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	var Photo entity.Photo

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email")
	}).First(&Photo, "ID = ?", c.Param("photoID")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func (s *PhotoHandlerImpl) PhotoDelete(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	photo := entity.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.ID = uint(photoID)

	err := db.Model(&photo).Where("id = ?", photoID).Delete(&photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
