package router

import (
	"finalProject2/controller"
	"finalProject2/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controller.NewUserHandlerImpl().UserRegister)

		userRouter.POST("/login", controller.NewUserHandlerImpl().UserLogin)

		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/:userID", middleware.UserAuthorization(), controller.NewUserHandlerImpl().UserUpdate)

		userRouter.DELETE("/:userID", middleware.UserAuthorization(), controller.NewUserHandlerImpl().UserDelete)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/create", controller.NewPhotoHandlerImpl().PhotoCreate)

		photoRouter.GET("/get", controller.NewPhotoHandlerImpl().PhotoGet)
		photoRouter.GET("/get/:photoID", controller.NewPhotoHandlerImpl().PhotoGetOne)

		photoRouter.DELETE("/delete/:photoID", middleware.PhotoAuthorization(), controller.NewPhotoHandlerImpl().PhotoDelete)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/create", middleware.CommentAuthentication(), middleware.CommentCreateAuthorization(), controller.NewCommentHandlerImpl().CommentCreate)
		commentRouter.GET("/get", controller.NewCommentHandlerImpl().CommentGet)
		commentRouter.PUT("/update/:commentID", middleware.CommentAuthorization(), controller.NewCommentHandlerImpl().CommentUpdate)
		commentRouter.DELETE("/delete/:commentID", middleware.CommentAuthorization(), controller.NewCommentHandlerImpl().CommentDelete)
	}

	socialmediaRouter := r.Group("/socialmedia")
	{
		socialmediaRouter.Use(middleware.Authentication())
		socialmediaRouter.POST("/create", controller.NewSocialMediaHandlerImpl().SocialMediaCreate)
		socialmediaRouter.GET("/get", controller.NewSocialMediaHandlerImpl().SocialMediaGet)
		socialmediaRouter.PUT("/update/:socialmediaID", middleware.SocialMediaAuthorization(), controller.NewSocialMediaHandlerImpl().SocialMediaUpdate)
		socialmediaRouter.DELETE("/delete/:socialmediaID", middleware.SocialMediaAuthorization(), controller.NewSocialMediaHandlerImpl().SocialMediaDelete)
	}

	return r
}
