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

type SocialMediaHandlerImpl struct{}

func NewSocialMediaHandlerImpl() SocialMediaHandler {
	return &SocialMediaHandlerImpl{}
}

func (s *SocialMediaHandlerImpl) SocialMediaCreate(c *gin.Context) {
	var db = database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	userID := uint(userData["id"].(float64))

	SocialMedia := entity.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
		"user_id":          userID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func (s *SocialMediaHandlerImpl) SocialMediaGet(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	var SocialMedia []entity.SocialMedia

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username")
	}).Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func (s *SocialMediaHandlerImpl) SocialMediaUpdate(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	SocialMedia := entity.SocialMedia{}

	socialmediaID, _ := strconv.Atoi(c.Param("socialmediaID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.ID = uint(socialmediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaID).Updates(map[string]interface{}{
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
		"user_id":          userID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func (s *SocialMediaHandlerImpl) SocialMediaDelete(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	SocialMedia := entity.SocialMedia{}

	socialmediaID, _ := strconv.Atoi(c.Param("socialmediaID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.ID = uint(socialmediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialmediaID).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
