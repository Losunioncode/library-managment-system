package handlers_page

import (
	"github.com/gin-gonic/gin"

	"github/losunioncode/library-managment-system/internal/server/controllers"
)

func InitRoutes(server *gin.Engine) {
	page := server.Group("/page")
	{
		page.GET("/search", controllers.HandleSearchPage)
		page.GET("/searchTitle", controllers.HandleSearchTitlePage)
		page.GET("/searchISBN", controllers.HandleSearchISBNPage)
		page.GET("/user/create_new", controllers.HandlerCreateUserPage)
		page.GET("/user/login", controllers.HandlerLoginUserPage)

	}
}
