package controller

import "github.com/gin-gonic/gin"

type SocialMediaHandler interface {
	SocialMediaCreate(*gin.Context)
	SocialMediaGet(*gin.Context)
	SocialMediaUpdate(*gin.Context)
	SocialMediaDelete(*gin.Context)
}
