package handlers_book

import (
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/middlewares"
	"github/losunioncode/library-managment-system/internal/server/controllers"
)

func InitRoutes(server *gin.Engine) {
	booklist := server.Group("/booklist")
	{
		booklist.POST("/searchByISBN", controllers.SearchISBNBook)
		booklist.POST("/searchByAuthor", controllers.SearchBookAuthor)
		booklist.POST("/searchByTitle", controllers.SearchBookTitle)

		booklist.GET("/getHeader", controllers.HandleHeader)
		secured := booklist.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/register", controllers.HandleRegisterNewBook)
		}
	}
}
