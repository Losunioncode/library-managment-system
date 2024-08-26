package server

import (
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/database"
	"github/losunioncode/library-managment-system/internal/server/controllers/routes/handlers-book"
	handlers_page "github/losunioncode/library-managment-system/internal/server/controllers/routes/handlers-page"
	handlers_user "github/losunioncode/library-managment-system/internal/server/controllers/routes/handlers-user"
	"os"
	"path/filepath"
)

//var templates = template.Must(template.ParseGlob("tmpl/*"))

var curr_dir, _ = os.Getwd()

func InitServer() {
	database.InitDB()

	server := gin.Default()
	server.LoadHTMLGlob(filepath.Join(curr_dir, "/internal/server/tmpl/**/*"))

	handlers_book.InitRoutes(server)
	handlers_page.InitRoutes(server)
	handlers_user.InitRoutes(server)

	server.Run(":8080")
}
