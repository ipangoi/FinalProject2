package middleware

import (
	"bytes"
	"encoding/json"
	"finalProject2/helper"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}

type CommentAuthenticationData struct {
	PhotoID float64 `json:"photo_id"`
}

func CommentAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawJSON, err := c.GetRawData()
		fmt.Println("Raw JSON request:", string(rawJSON))
		fmt.Println("Request Headers:", c.Request.Header)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Error reading raw JSON data",
			})
			c.Abort()
			return
		}

		var requestBody map[string]interface{}
		if err := json.Unmarshal(rawJSON, &requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid JSON payload",
			})
			c.Abort()
			return
		}

		photoID, ok := requestBody["photo_id"].(float64)
		if !ok || photoID != float64(uint64(photoID)) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "'photo_id' must be a non-negative integer",
			})
			c.Abort()
			return
		}

		c.Set("photoData", map[string]interface{}{"id": uint(photoID)})

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawJSON))

		c.Next()
	}
}
