package handlers_page

import (
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/middlewares"
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
		secured := page.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/borrow", controllers.HandleRegisterNewBook)
			secured.GET("/passwordToChange", controllers.HandlePasswordToChangePage)
			secured.GET("/extend", controllers.HandleExtendDeadlinePage)
			secured.GET("/return", controllers.HandleReturnBookPage)
			secured.GET("/checkDeadline", controllers.HandleBorrowBookPage)
		}
	}
}
