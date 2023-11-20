package controller

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	CommentCreate(*gin.Context)
	CommentGet(*gin.Context)
	CommentUpdate(*gin.Context)
	CommentDelete(*gin.Context)
}
