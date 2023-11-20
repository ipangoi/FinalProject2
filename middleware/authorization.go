package middleware

import (
	"finalProject2/database"
	"finalProject2/entity"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userID, err := strconv.Atoi(c.Param("userID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		user := entity.User{}

		err = db.Select("id").First(&user, uint(userID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesnt exist",
			})
			return
		}

	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoID, err := strconv.Atoi(c.Param("photoID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := entity.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesnt exist",
			})
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		}

		c.Next()

	}
}

func CommentCreateAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photoData := c.MustGet("photoData").(map[string]interface{})
		photoID := uint(photoData["id"].(uint))
		var photo entity.Photo
		if err := database.GetDB().Where("id = ?", photoID).First(&photo).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Photo not found",
			})
			return
		}

		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()

	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentID, err := strconv.Atoi(c.Param("commentID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := entity.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesnt exist",
			})
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		}

		c.Next()

	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialmediaID, err := strconv.Atoi(c.Param("socialmediaID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Social_Media := entity.SocialMedia{}

		err = db.Select("user_id").First(&Social_Media, uint(socialmediaID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesnt exist",
			})
			return
		}

		if Social_Media.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
		}

		c.Next()

	}
}
