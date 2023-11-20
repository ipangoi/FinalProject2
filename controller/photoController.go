package controller

import "github.com/gin-gonic/gin"

type PhotoHandler interface {
	PhotoCreate(*gin.Context)
	PhotoGet(*gin.Context)
	PhotoGetOne(*gin.Context)
	PhotoDelete(*gin.Context)
}
