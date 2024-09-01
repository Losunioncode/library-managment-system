package handlers_user

import (
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/middlewares"
	"github/losunioncode/library-managment-system/internal/server/controllers"
)

func InitRoutes(server *gin.Engine) {
	api := server.Group("/api")
	{
		api.POST("/registerUser", controllers.RegisterUser)
		api.POST("/user/login", controllers.GenerateToken)
		api.GET("/user/logout", controllers.HandleLogoutUser)
		api.POST("/user/registerUser", controllers.RegisterUser)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.POST("/changePasswordUser", controllers.PasswordToChangeHandler)
		}
	}
}
