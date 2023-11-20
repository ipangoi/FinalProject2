package controller

import (
	"encoding/json"
	"finalProject2/database"
	"finalProject2/entity"
	"finalProject2/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandlerImpl struct{}

func NewCommentHandlerImpl() CommentHandler {
	return &CommentHandlerImpl{}
}

func (s *CommentHandlerImpl) CommentCreate(c *gin.Context) {
	var db = database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoData := c.MustGet("photoData").(map[string]interface{})
	contentType := helper.GetContentType(c)

	userID := uint(userData["id"].(float64))
	photoID := uint(photoData["id"].(uint))
	Comment := entity.Comment{}

	rawJSON, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Error reading raw JSON data",
		})
		return
	}

	fmt.Println("Raw JSON request:", string(rawJSON))
	fmt.Println("Request Headers:", c.Request.Header)

	if contentType == appJSON {
		if err := json.Unmarshal(rawJSON, &Comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid JSON payload for Comment",
			})
			return
		}
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = photoID

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func (s *CommentHandlerImpl) CommentGet(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	var Comment []entity.Comment

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email")
	}).Preload("Photo").Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func (s *CommentHandlerImpl) CommentUpdate(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	Comment := entity.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(commentID)

	err := db.Model(&Comment).Where("id = ?", commentID).Updates(map[string]interface{}{
		"message": Comment.Message,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.Preload("Photo").First(&Comment, Comment.ID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	fmt.Println(Comment.PhotoID)
	fmt.Println(Comment.Photo.ID)
	if Comment.Photo.ID != 0 && Comment.PhotoID == Comment.Photo.ID {
		c.JSON(http.StatusOK, gin.H{
			"id":         Comment.ID,
			"title":      Comment.Photo.Title,
			"caption":    Comment.Photo.Caption,
			"message":    Comment.Message,
			"photo_url":  Comment.Photo.PhotoURL,
			"user_id":    userID,
			"updated_at": Comment.UpdatedAt,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Comment and Photo IDs do not match",
		})
	}
}

func (s *CommentHandlerImpl) CommentDelete(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	Comment := entity.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.ID = uint(commentID)

	err := db.Model(&Comment).Where("id = ?", commentID).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
